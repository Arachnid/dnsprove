// Copyright 2017 Nick Johnson. All rights reserved.
// Portions copyright 2011 Miek Gieben.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// dnsprove is a utility that submits DNSSEC signatures to an Ethereum oracle,
// allowing you to prove the (non)existence and contents of DNS records onchain.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/miekg/dns"
)

var (
	hashes          = flag.String("hashes", "SHA256", "a comma-separated list of supported hash algorithms")
	algorithms      = flag.String("algorithms", "RSASHA256", "a comma-separated list of supported digest algorithms")
	trustAnchors = []*dns.DS{
		&dns.DS{
			Hdr: dns.RR_Header{Name: ".", Rrtype: dns.TypeDS, Class: dns.ClassINET},
			KeyTag: 19036,
			Algorithm: 8,
			DigestType: 2,
			Digest: "49AAC11D7B6F6446702E54A1607371607A1A41855200FD2CE1CDDE32F24E8FB5",
		},
		&dns.DS{
			Hdr: dns.RR_Header{Name: ".", Rrtype: dns.TypeDS, Class: dns.ClassINET},
			KeyTag: 20326,
			Algorithm: 8,
			DigestType: 2,
			Digest: "E06D44B80B8F1D39A95C0B0D7C65D08458E880409BBC683457104237C7F8EC8D",
		},
	}
)

type signedSet struct {
	sig *dns.RRSIG
	rrs []dns.RR
}

type dnskeyEntry struct {
	name string
	algorithm uint8
	keytag uint16
}

type Client struct {
	c *dns.Client
	nameserver string
	knownHashes map[dnskeyEntry][]*dns.DS
	supportedAlgorithms map[uint8]struct{}
	supportedDigests map[uint8]struct{}
}

func (client *Client) addDS(ds *dns.DS) {
	key := dnskeyEntry{ds.Header().Name, ds.Algorithm, ds.KeyTag}
	client.knownHashes[key] = append(client.knownHashes[key], ds)
}

func (client *Client) supportsAlgorithm(algorithm uint8) bool {
	_, ok := client.supportedAlgorithms[algorithm]
	return ok
}

func (client *Client) supportsDigest(digest uint8) bool {
	_, ok := client.supportedDigests[digest]
	return ok
}

func NewClient(nameserver string, roots []*dns.DS, algorithms, digests map[uint8]struct{}) *Client {
	client := &Client{
		c: new(dns.Client),
		nameserver: nameserver,
		knownHashes: make(map[dnskeyEntry][]*dns.DS),
		supportedAlgorithms: algorithms,
		supportedDigests: digests,
	}
	for _, root := range roots {
		client.addDS(root)
	}
	return client
}

func (client *Client) Query(qtype uint16, qclass uint16, name string) (*dns.Msg, error) {
	m := &dns.Msg{
		MsgHdr: dns.MsgHdr{
			Authoritative:     false,
			AuthenticatedData: false,
			CheckingDisabled:  false,
			RecursionDesired:  true,
			Opcode:            dns.OpcodeQuery,
		},
		Question: []dns.Question{
			dns.Question{
				Name: dns.Fqdn(name),
				Qtype: qtype,
				Qclass: qclass,
			},
		},
	}

	o := &dns.OPT{
		Hdr: dns.RR_Header{
			Name:   ".",
			Rrtype: dns.TypeOPT,
		},
	}
	o.SetDo()
	o.SetUDPSize(dns.DefaultMsgSize)
	m.Extra = append(m.Extra, o)
	m.Id = dns.Id()

	r, _, err := client.c.Exchange(m, client.nameserver)
	log.Printf("%s\n", r.String())
	return r, err
}

func (client *Client) QueryWithProof(qtype, qclass uint16, name string) ([]signedSet, error) {
		if name[len(name) - 1] != '.' {
			name = name + "."
		}

		r, err := client.Query(qtype, qclass, name)
		if err != nil {
			return nil, err
		}

		sigs := filterRRs(r.Answer, dns.TypeRRSIG)
		rrs := getRRset(r.Answer, name, qtype)
		if len(sigs) == 0 || len(rrs) == 0 {
			return nil, fmt.Errorf("No signed RRSETs available for %s %s %s", dns.ClassToString[qclass], dns.TypeToString[qtype], name)
		}

		for _, sig := range sigs {
			sig := sig.(*dns.RRSIG)
			ret, err := client.verifyRRSet(sig, rrs)
			if err == nil {
				ret = append(ret, signedSet{sig, rrs})
				return ret, nil
			}
			log.Printf("Failed to verify RRSET for %s %s %s with signature %s (%s/%d): %s\n", dns.ClassToString[qclass], dns.TypeToString[qtype], name, sig.SignerName, dns.AlgorithmToString[sig.Algorithm], sig.KeyTag, err)
		}

		return nil, fmt.Errorf("Could not validate %s %s %s: no valid signatures found", dns.ClassToString[qclass], dns.TypeToString[qtype], name)
}

func (client *Client) verifyRRSet(sig *dns.RRSIG, rrs []dns.RR) ([]signedSet, error) {
	if !client.supportsAlgorithm(sig.Algorithm) {
		return nil, fmt.Errorf("Unsupported algorithm: %s", dns.AlgorithmToString[sig.Algorithm])
	}

	var sets []signedSet
	var keys []dns.RR
	var err error
	if sig.Header().Name == sig.SignerName && rrs[0].Header().Rrtype == dns.TypeDNSKEY {
		// RRSet is self-signed; verify against itself
		keys = rrs
	} else {
		// Find the keys that signed this RRSET
		sets, err = client.QueryWithProof(dns.TypeDNSKEY, sig.Header().Class, sig.SignerName)
		keys = sets[len(sets)-1].rrs
	}
	if err != nil {
		return nil, err
	}

	// Iterate over the keys looking for one that validly signs our RRSET
	for _, key := range keys {
		key := key.(*dns.DNSKEY)
		if key.Algorithm != sig.Algorithm || key.KeyTag() != sig.KeyTag || key.Header().Name != sig.SignerName {
			continue
		}
		if err := sig.Verify(key, rrs); err != nil {
			fmt.Printf("Could not verify signature on %s %s %s with %s (%s/%d): %s", dns.ClassToString[sig.Header().Class], dns.TypeToString[sig.Header().Rrtype], sig.Header().Name, key.Header().Name, dns.AlgorithmToString[key.Algorithm], key.KeyTag(), err)
		}
		if sig.Header().Name == sig.SignerName {
			// RRSet is self-signed; look for DS records in parent zones to verify
			sets, err = client.verifyWithDS(key)
			if err != nil {
				return nil, err
			}
		}
		return sets, nil
	}
	return nil, fmt.Errorf("Could not validate signature for %s %s %s (%s/%d); no valid keys found", dns.ClassToString[sig.Header().Class], dns.TypeToString[sig.Header().Rrtype], sig.Header().Name, dns.AlgorithmToString[sig.Algorithm], sig.KeyTag)
}

func (client *Client) verifyWithDS(key *dns.DNSKEY) ([]signedSet, error) {
	keytag := key.KeyTag()
	// Check the roots
	for _, ds := range client.knownHashes[dnskeyEntry{key.Header().Name, key.Algorithm, keytag}] {
		if !client.supportsDigest(ds.DigestType) {
			continue
		}
		if strings.ToLower(key.ToDS(ds.DigestType).Digest) == strings.ToLower(ds.Digest) {
			return []signedSet{}, nil
		}
	}

	// Look up the DS record
	sets, err := client.QueryWithProof(dns.TypeDS, key.Header().Class, key.Header().Name)
	if err != nil {
		return nil, err
	}
	for _, ds := range sets[len(sets) - 1].rrs {
		ds := ds.(*dns.DS)
		if !client.supportsDigest(ds.DigestType) {
			continue
		}
		if strings.ToLower(key.ToDS(ds.DigestType).Digest) == strings.ToLower(ds.Digest) {
			return sets, nil
		}
	}
	return nil, fmt.Errorf("Could not find any DS records that validate %s DNSKEY %s (%s/%d)", dns.ClassToString[key.Header().Class], key.Header().Name, dns.AlgorithmToString[key.Algorithm], keytag)
}

func filterRRs(rrs []dns.RR, qtype uint16) []dns.RR {
	ret := make([]dns.RR, 0)
	for _, rr := range rrs {
		if rr.Header().Rrtype == qtype {
			ret = append(ret, rr)
		}
	}
	return ret
}

func getRRset(rrs []dns.RR, name string, qtype uint16) []dns.RR {
	var ret []dns.RR
	for _, rr := range rrs {
		if strings.ToLower(rr.Header().Name) == strings.ToLower(name) && rr.Header().Rrtype == qtype {
			ret = append(ret, rr)
		}
	}
	return ret
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] qtype name\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		return
	}

	qtype, ok := dns.StringToType[flag.Arg(0)]
	if !ok {
		log.Fatalf("Unrecognised query type: %s\n", flag.Arg(0))
	}
	qclass := uint16(dns.ClassINET)
	name := flag.Arg(1)

	hashmap := make(map[uint8]struct{})
	for _, hashname := range strings.Split(*hashes, ",") {
		hashmap[dns.StringToHash[hashname]] = struct{}{}
	}

	algmap := make(map[uint8]struct{})
	for _, algname := range strings.Split(*algorithms, ",") {
		algmap[dns.StringToAlgorithm[algname]] = struct{}{}
	}

	conf, err := dns.ClientConfigFromFile("/etc/resolv.conf")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	nameserver := conf.Servers[0] + ":53"

	client := NewClient(nameserver, trustAnchors, algmap, hashmap)
	proofs, err := client.QueryWithProof(qtype, qclass, name)
	if err != nil {
		log.Fatalf("Error resolving: %v\n", err)
		return
	}

	fmt.Printf("Proofs:\n")
	for _, proof := range proofs {
		fmt.Printf("\n%s\n", proof.sig.String())
		for _, rr := range proof.rrs {
			fmt.Printf("%s\n", rr.String())
		}
	}
}
