// Copyright 2019 Nick Johnson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// dnsprove is a utility that submits DNSSEC signatures to an Ethereum oracle,
// allowing you to prove the (non)existence and contents of DNS records onchain.
package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/arachnid/dnsprove/ens"
	"github.com/arachnid/dnsprove/oracle"
	"github.com/arachnid/dnsprove/proofs"
	"github.com/arachnid/dnsprove/registrar"
	"github.com/arachnid/dnsprove/root"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/inconshreveable/log15"
	"github.com/miekg/dns"
	prompt "github.com/segmentio/go-prompt"
)

var (
	server     = flag.String("server", "https://dns.google.com/experimental", "The URL of the dns-over-https server to use")
	hashes     = flag.String("hashes", "SHA1,SHA256", "a comma-separated list of supported hash algorithms")
	algorithms = flag.String("algorithms", "RSASHA1,RSASHA1-NSEC3-SHA1,RSASHA256", "a comma-separated list of supported digest algorithms")
	verbosity  = flag.Int("verbosity", 3, "logging level verbosity (0-4)")
	rpc        = flag.String("rpc", "http://localhost:8545", "RPC path to Ethereum node")
	keyfile    = flag.String("keyfile", "", "Path to JSON keyfile")
	insecure   = flag.Bool("insecure", false, "Do not prompt for a password, assume the empty string")
	gasprice   = flag.Float64("gasprice", 5.0, "Gas price, in gwei")

	proveFlags    = flag.NewFlagSet("prove", flag.ExitOnError)
	oracleAddress = proveFlags.String("address", "", "Contract address for DNSSEC oracle")
	print         = proveFlags.Bool("print", false, "don't upload to the contract, just print proof data")
	yes           = proveFlags.Bool("yes", false, "Do not prompt before sending transactions")

	claimFlags      = flag.NewFlagSet("claim", flag.ExitOnError)
	registryAddress = claimFlags.String("address", "0x314159265dd8dbb310642f98f50c066173c1259b", "Contract address for ENS registry")

	subcommands = map[string]func([]string){
		"prove": proveCommand,
		"claim": claimCommand,
	}

	trustAnchors = []*dns.DS{
		&dns.DS{
			Hdr:        dns.RR_Header{Name: ".", Rrtype: dns.TypeDS, Class: dns.ClassINET},
			KeyTag:     20326,
			Algorithm:  8,
			DigestType: 2,
			Digest:     "E06D44B80B8F1D39A95C0B0D7C65D08458E880409BBC683457104237C7F8EC8D",
		},
	}

	NotDNSSECEnabledError = errors.New("RR does not exist and either no NSEC records returned or NSEC records are unsigned")
)

type dnskeyEntry struct {
	name      string
	algorithm uint8
	keytag    uint16
}

type Client struct {
	c                   *dns.Client
	nameserver          string
	knownHashes         map[dnskeyEntry][]*dns.DS
	supportedAlgorithms map[uint8]struct{}
	supportedDigests    map[uint8]struct{}
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
		c:                   new(dns.Client),
		nameserver:          nameserver,
		knownHashes:         make(map[dnskeyEntry][]*dns.DS),
		supportedAlgorithms: algorithms,
		supportedDigests:    digests,
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
				Name:   dns.Fqdn(name),
				Qtype:  qtype,
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

	req, err := m.Pack()
	if err != nil {
		return nil, err
	}

	response, err := http.Post(client.nameserver, "application/dns-udpwireformat", bytes.NewReader(req))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Got unexpected status from server: %s", response.Status)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var r dns.Msg
	err = r.Unpack(data)
	if err == nil {
		log.Debug("DNS response:\n" + r.String())
		log.Info("DNS query", "class", dns.ClassToString[qclass], "type", dns.TypeToString[qtype], "name", name, "answer", len(r.Answer), "extra", len(r.Extra), "ns", len(r.Ns))
	}
	return &r, err
}

func (client *Client) QueryWithProof(qtype, qclass uint16, name string) ([]proofs.SignedSet, bool, error) {
	found := false

	if name[len(name)-1] != '.' {
		name = name + "."
	}

	r, err := client.Query(qtype, qclass, name)
	if err != nil {
		return nil, false, err
	}

	rrs := getRRset(r.Answer, name, qtype)
	var sigs []dns.RR
	if len(rrs) > 0 {
		found = true
		sigs = findSignatures(r.Answer, name)
		if len(sigs) == 0 {
			return nil, false, fmt.Errorf("No signed RRSETs available for %s %s", dns.TypeToString[qtype], name)
		}
	} else {
		rrs = getNSECRRs(r.Ns, name)
		if len(rrs) == 0 {
			return nil, false, NotDNSSECEnabledError
		}
		log.Info("RR does not exist; got NSEC", "qtype", dns.TypeToString[qtype], "name", name)
		sigs = findSignatures(r.Ns, rrs[0].Header().Name)
		if len(sigs) == 0 {
			return nil, false, NotDNSSECEnabledError
		}
	}

	for _, sig := range sigs {
		sig := sig.(*dns.RRSIG)
		if sig.TypeCovered != rrs[0].Header().Rrtype {
			continue
		}
		ret, err := client.verifyRRSet(sig, rrs)
		if err == nil {
			result := proofs.SignedSet{sig, rrs, name}
			ret = append(ret, result)
			return ret, found, nil
		}
		log.Warn("Failed to verify RRSET", "type", dns.TypeToString[rrs[0].Header().Rrtype], "name", name, "signername", sig.SignerName, "algorithm", dns.AlgorithmToString[sig.Algorithm], "keytag", sig.KeyTag, "err", err)
	}

	return nil, found, fmt.Errorf("Could not validate %s %s %s: no valid signatures found", dns.ClassToString[qclass], dns.TypeToString[qtype], name)
}

func (client *Client) verifyRRSet(sig *dns.RRSIG, rrs []dns.RR) ([]proofs.SignedSet, error) {
	if !client.supportsAlgorithm(sig.Algorithm) {
		return nil, fmt.Errorf("Unsupported algorithm: %s", dns.AlgorithmToString[sig.Algorithm])
	}

	var sets []proofs.SignedSet
	var keys []dns.RR
	var err error
	if sig.Header().Name == sig.SignerName && rrs[0].Header().Rrtype == dns.TypeDNSKEY {
		// RRSet is self-signed; verify against itself
		keys = rrs
	} else {
		// Find the keys that signed this RRSET
		var found bool
		sets, found, err = client.QueryWithProof(dns.TypeDNSKEY, sig.Header().Class, sig.SignerName)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, fmt.Errorf("DNSKEY %s not found", sig.SignerName)
		}
		keys = sets[len(sets)-1].Rrs
	}

	// Iterate over the keys looking for one that validly signs our RRSET
	for _, key := range keys {
		key := key.(*dns.DNSKEY)
		if key.Algorithm != sig.Algorithm || key.KeyTag() != sig.KeyTag || key.Header().Name != sig.SignerName {
			continue
		}
		if err := sig.Verify(key, rrs); err != nil {
			log.Error("Could not verify signature", "type", dns.TypeToString[rrs[0].Header().Rrtype], "signame", sig.Header().Name, "keyname", key.Header().Name, "algorithm", dns.AlgorithmToString[key.Algorithm], "keytag", key.KeyTag(), "key", key, "rrs", rrs, "sig", sig, "err", err)
			continue
		}
		if sig.Header().Name == sig.SignerName && rrs[0].Header().Rrtype == dns.TypeDNSKEY {
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

func (client *Client) verifyWithDS(key *dns.DNSKEY) ([]proofs.SignedSet, error) {
	keytag := key.KeyTag()
	// Check the roots
	for _, ds := range client.knownHashes[dnskeyEntry{key.Header().Name, key.Algorithm, keytag}] {
		if !client.supportsDigest(ds.DigestType) {
			continue
		}
		if strings.ToLower(key.ToDS(ds.DigestType).Digest) == strings.ToLower(ds.Digest) {
			return []proofs.SignedSet{}, nil
		}
	}

	// If it's a root DS, and we don't have it in our roots, no point querying for it.
	if key.Header().Name == "." {
		return nil, fmt.Errorf("DS . with key tag %d not found", keytag)
	}

	// Look up the DS record
	sets, found, err := client.QueryWithProof(dns.TypeDS, key.Header().Class, key.Header().Name)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("DS %s not found", key.Header().Name)
	}
	for _, ds := range sets[len(sets)-1].Rrs {
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

func findSignatures(rrs []dns.RR, name string) []dns.RR {
	ret := make([]dns.RR, 0)
	for _, rr := range rrs {
		// TODO: Wildcard support
		if rr.Header().Rrtype == dns.TypeRRSIG && rr.Header().Name == name {
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

func getNSECRRs(rrs []dns.RR, name string) []dns.RR {
	ret := make([]dns.RR, 0)
	for _, rr := range rrs {
		if nsec, ok := rr.(*dns.NSEC); ok && nsecCovers(rr.Header().Name, name, nsec.NextDomain) {
			ret = append(ret, rr)
		}
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func compareDomainNames(a, b string) int {
	alabels := dns.SplitDomainName(a)
	blabels := dns.SplitDomainName(b)

	for i := 1; i <= min(len(alabels), len(blabels)); i++ {
		result := strings.Compare(alabels[len(alabels)-i], blabels[len(blabels)-i])
		if result != 0 {
			return result
		}
	}

	return len(alabels) - len(blabels)
}

func nsecCovers(owner, test, next string) bool {
	owner = strings.ToLower(owner)
	test = strings.ToLower(test)
	next = strings.ToLower(next)
	return compareDomainNames(owner, test) <= 0 && (compareDomainNames(test, next) <= 0 || strings.HasSuffix(test, next))
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] command\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*verbosity), log.StreamHandler(os.Stderr, log.TerminalFormat())))

	subcommand, ok := subcommands[flag.Arg(0)]
	if !ok {
		flag.Usage()
		return
	}

	subcommand(flag.Args()[1:])
}

func proveCommand(args []string) {
	proveFlags.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] prove [prove options] qtype qname\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nGeneral options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nProve command options:\n")
		proveFlags.PrintDefaults()
	}
	proveFlags.Parse(args)

	if proveFlags.NArg() != 2 {
		proveFlags.Usage()
		return
	}

	qtype, ok := dns.StringToType[proveFlags.Arg(0)]
	if !ok {
		log.Crit("Unrecognised query type", "qtype", qtype)
		os.Exit(1)
	}
	name := proveFlags.Arg(1)

	sets, found, err := getProofs(qtype, name)
	if err != nil {
		log.Crit("Error resolving", "qtype", qtype, "name", name, "err", err)
		os.Exit(1)
	}

	if *print {
		for _, proof := range sets {
			fmt.Printf("\n// %s\n", proof.Sig.String())
			for _, rr := range proof.Rrs {
				for _, line := range strings.Split(rr.String(), "\n") {
					fmt.Printf("// %s\n", line)
				}
			}
			data, err := proof.Pack()
			if err != nil {
				log.Crit("Error packing RRSet", "err", err)
				os.Exit(1)
			}
			sig, err := proof.PackSignature()
			if err != nil {
				log.Crit("Error packing RRSet signature", "err", err)
				os.Exit(1)
			}
			fmt.Printf("[\"%s\", \"%x\", \"%x\"],\n", proof.Name, data, sig)
		}
		os.Exit(0)
	}

	conn, err := ethclient.Dial(*rpc)
	if err != nil {
		log.Crit("Error connecting to Ethereum node", "err", err)
		os.Exit(1)
	}

	o, err := oracle.New(common.HexToAddress(*oracleAddress), conn)
	if err != nil {
		log.Crit("Error creating oracle", "err", err)
		os.Exit(1)
	}

	if !found {
		// We're deleting a domain. If it's not already there, there's nothing to do.
		_, _, hash, err := o.Rrdata(qtype, name)
		if err != nil {
			log.Crit("Error checking RRDATA", "qtype", qtype, "name", name, "err", err)
			os.Exit(1)
		}
		if hash == [20]byte{} {
			fmt.Printf("RRSet not found in oracle. Nothing to do; exiting\n")
			os.Exit(0)
		}
	} else {
		// If the RRset already matches, there's nothing to do
		matches, err := o.RecordMatches(sets[len(sets)-1])
		if err != nil {
			log.Crit("Error checking for record", "err", err)
			os.Exit(1)
		}
		if matches && found {
			fmt.Printf("Nothing to do; exiting.\n")
			os.Exit(0)
		}
	}

	known, err := o.FindFirstUnknownProof(sets)
	if err != nil {
		log.Crit("Error checking proofs against oracle", "err", err)
		os.Exit(1)
	}

	if !*yes {
		if !prompt.Confirm("Send a transaction to prove %s %s (%d proofs) onchain?", dns.TypeToString[sets[len(sets)-1].Rrs[0].Header().Rrtype], name, len(sets)-known) {
			fmt.Printf("Exiting at user request.\n")
			return
		}
	}

	auth, err := makeTransactor(conn)
	if err != nil {
		log.Crit("Could not create transactor", "err", err)
		os.Exit(1)
	}

	var txs []*types.Transaction
	if found {
		tx, err := o.SendProofs(auth, sets, known)
		if err != nil {
			log.Crit("Error sending proofs", "err", err)
			os.Exit(1)
		}
		txs = append(txs, tx)
	} else {
		nsec := sets[len(sets)-1]
		if known < len(sets)-1 {
			tx, err := o.SendProofs(auth, sets[:len(sets)-1], known)
			if err != nil {
				log.Crit("Error sending proofs", "err", err)
				os.Exit(1)
			}
			txs = append(txs, tx)
			auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
		}

		proof, err := sets[len(sets)-2].PackRRSet()
		if err != nil {
			log.Crit("Could not pack proof for record", "err", err)
			os.Exit(1)
		}

		deletetx, err := o.DeleteRRSet(auth, qtype, name, nsec, proof)
		if err != nil {
			log.Crit("Error deleting RRSet", "err", err)
			os.Exit(1)
		}
		txs = append(txs, deletetx)
	}

	txids := make([]string, 0, len(txs))
	for _, tx := range txs {
		txids = append(txids, tx.Hash().String())
	}
	log.Info("Transactions sent", "txids", txids)
}

func makeTransactor(conn *ethclient.Client) (*bind.TransactOpts, error) {
	key, err := os.Open(*keyfile)
	if err != nil {
		log.Crit("Could not open keyfile", "err", err)
		os.Exit(1)
	}

	pass := ""
	if !*insecure {
		pass = prompt.Password("Password")
	}
	auth, err := bind.NewTransactor(key, pass)
	if err != nil {
		return nil, err
	}
	auth.GasPrice = big.NewInt(int64(*gasprice * 1000000000))
	if err = updateNonce(conn, auth); err != nil {
		return nil, err
	}
	return auth, nil
}

func updateNonce(conn *ethclient.Client, auth *bind.TransactOpts) error {
	nonce, err := conn.PendingNonceAt(context.TODO(), auth.From)
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	return nil
}

func getProofs(qtype uint16, name string) ([]proofs.SignedSet, bool, error) {
	qclass := uint16(dns.ClassINET)
	if !strings.HasSuffix(name, ".") {
		name = name + "."
	}

	hashmap := make(map[uint8]struct{})
	for _, hashname := range strings.Split(*hashes, ",") {
		hashmap[dns.StringToHash[hashname]] = struct{}{}
	}

	algmap := make(map[uint8]struct{})
	for _, algname := range strings.Split(*algorithms, ",") {
		algmap[dns.StringToAlgorithm[algname]] = struct{}{}
	}

	client := NewClient(*server, trustAnchors, algmap, hashmap)
	return client.QueryWithProof(qtype, qclass, name)
}

func claimCommand(args []string) {
	claimFlags.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] claim [claim options] name\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nGeneral options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nClaim command options:\n")
		claimFlags.PrintDefaults()
	}
	claimFlags.Parse(args)

	if claimFlags.NArg() != 1 {
		claimFlags.Usage()
		return
	}

	conn, err := ethclient.Dial(*rpc)
	if err != nil {
		log.Crit("Error connecting to Ethereum node", "err", err)
		os.Exit(1)
	}

	name := claimFlags.Arg(0)

	registry, err := ens.New(common.HexToAddress(*registryAddress), conn)
	if err != nil {
		log.Crit("Error instantiating registry", "err", err)
		os.Exit(1)
	}

	parentName := ""
	if nameparts := strings.SplitN(name, ".", 2); len(nameparts) > 1 {
		parentName = nameparts[1]
	}

	addr, err := registry.Owner(parentName)
	if err != nil {
		log.Crit("Could not get ENS owner", "name", parentName, "err", err)
		os.Exit(1)
	}

	if addr == (common.Address{}) {
		log.Crit("Name does not exist in ENS", "name", parentName)
		os.Exit(1)
	}

	reg, err := registrar.New(addr, conn)
	if err == nil {
		if err := claimWithRegistrar(conn, name, reg); err != nil {
			log.Crit("Error claiming name with registrar", "name", name, "error", err)
			os.Exit(1)
		}
		return
	} else if err != registrar.InterfaceNotSupportedError {
		log.Crit("Could not instantiate DNSSEC registrar", "name", parentName, "err", err)
		os.Exit(1)
	}

	root, err := root.New(addr, conn)
	if err != nil {
		log.Crit("Could not instantiate root contract", "name", parentName, "err", err)
		os.Exit(1)
	}

	if err := claimWithRoot(conn, name, root); err != nil {
		log.Crit("Error claiming name with root contract", "name", name, "error", err)
		os.Exit(1)
	}
}

func claimWithRegistrar(conn *ethclient.Client, name string, registrar *registrar.DNSRegistrar) error {
	sets, found, err := getProofs(dns.TypeTXT, "_ens."+name)
	if err != nil {
		return err
	}

	auth, err := makeTransactor(conn)
	if err != nil {
		log.Crit("Could not create transactor", "err", err)
		os.Exit(1)
	}

	if found {
		tx, err := registrar.Claim(auth, name, sets)
		if err != nil {
			return err
		}
		log.Info("Sent transaction", "tx", tx.Hash().String())
	} else {
		txs, err := registrar.Unclaim(auth, name, sets)
		if err != nil {
			return err
		}
		txids := make([]string, 0, len(txs))
		for _, tx := range txs {
			txids = append(txids, tx.Hash().String())
		}
		log.Info("Transactions sent", "txids", txids)
	}

	return nil
}

func claimWithRoot(conn *ethclient.Client, name string, root *root.Root) error {
	sets, found, err := getProofs(dns.TypeTXT, "_ens.nic."+name)
	if err != nil && err != NotDNSSECEnabledError {
		return err
	}

	if found {
		auth, err := makeTransactor(conn)
		if err != nil {
			log.Crit("Could not create transactor", "err", err)
			os.Exit(1)
		}

		tx, err := root.Claim(auth, name, sets)
		if err != nil {
			return err
		}
		log.Info("Sent transaction", "tx", tx.Hash().String())
	} else {
		dssets, found, err := getProofs(dns.TypeDS, name)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("Cannot claim name %s: Not found in DNS", name)
		}

		auth, err := makeTransactor(conn)
		if err != nil {
			log.Crit("Could not create transactor", "err", err)
			os.Exit(1)
		}

		txs, err := root.ClaimDefault(auth, name, sets, dssets)
		if err != nil {
			return err
		}

		txids := make([]string, 0, len(txs))
		for _, tx := range txs {
			txids = append(txids, tx.Hash().String())
		}
		log.Info("Transactions sent", "txids", txids)
	}

	return nil
}
