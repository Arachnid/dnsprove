// Copyright 2011 Miek Gieben. All rights reserved.
// Portions copyright 2017 Nick Johnson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proofs

import (
  "bytes"
  "encoding/base64"
  "encoding/binary"
  "errors"
  "sort"
  "strings"

  "github.com/miekg/dns"
)

type SignedSet struct {
	Sig *dns.RRSIG
	Rrs []dns.RR
}

type wireSlice [][]byte

func (p wireSlice) Len() int      { return len(p) }
func (p wireSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p wireSlice) Less(i, j int) bool {
	_, ioff, _ := dns.UnpackDomainName(p[i], 0)
	_, joff, _ := dns.UnpackDomainName(p[j], 0)
	return bytes.Compare(p[i][ioff+10:], p[j][joff+10:]) < 0
}

func (ss *SignedSet) Pack() (buf []byte, sig []byte, err error) {
  sigwire := make([]byte, dns.Len(ss.Sig))
  off, err := packSigWire(ss.Sig, sigwire)
  if err != nil {
    return nil, nil, err
  }
  sigwire = sigwire[:off]

  rrdata, err := rawSignatureData(ss.Rrs, ss.Sig)
  if err != nil {
    return nil, nil, err
  }

  sigwire = append(sigwire, rrdata...)

  sig, err = base64.StdEncoding.DecodeString(ss.Sig.Signature)
  if err != nil {
    return nil, nil, err
  }

  return sigwire, sig, nil
}

// Return the raw signature data.
func rawSignatureData(rrset []dns.RR, s *dns.RRSIG) (buf []byte, err error) {
	wires := make(wireSlice, len(rrset))
	for i, r := range rrset {
		r1 := dns.Copy(r)
		r1.Header().Ttl = s.OrigTtl
		labels := dns.SplitDomainName(r1.Header().Name)
		// 6.2. Canonical RR Form. (4) - wildcards
		if len(labels) > int(s.Labels) {
			// Wildcard
			r1.Header().Name = "*." + strings.Join(labels[len(labels)-int(s.Labels):], ".") + "."
		}
		// RFC 4034: 6.2.  Canonical RR Form. (2) - domain name to lowercase
		r1.Header().Name = strings.ToLower(r1.Header().Name)
		// 6.2. Canonical RR Form. (3) - domain rdata to lowercase.
		//   NS, MD, MF, CNAME, SOA, MB, MG, MR, PTR,
		//   HINFO, MINFO, MX, RP, AFSDB, RT, SIG, PX, NXT, NAPTR, KX,
		//   SRV, DNAME, A6
		//
		// RFC 6840 - Clarifications and Implementation Notes for DNS Security (DNSSEC):
		//	Section 6.2 of [RFC4034] also erroneously lists HINFO as a record
		//	that needs conversion to lowercase, and twice at that.  Since HINFO
		//	records contain no domain names, they are not subject to case
		//	conversion.
		switch x := r1.(type) {
		case *dns.NS:
			x.Ns = strings.ToLower(x.Ns)
		case *dns.CNAME:
			x.Target = strings.ToLower(x.Target)
		case *dns.SOA:
			x.Ns = strings.ToLower(x.Ns)
			x.Mbox = strings.ToLower(x.Mbox)
		case *dns.MB:
			x.Mb = strings.ToLower(x.Mb)
		case *dns.MG:
			x.Mg = strings.ToLower(x.Mg)
		case *dns.MR:
			x.Mr = strings.ToLower(x.Mr)
		case *dns.PTR:
			x.Ptr = strings.ToLower(x.Ptr)
		case *dns.MINFO:
			x.Rmail = strings.ToLower(x.Rmail)
			x.Email = strings.ToLower(x.Email)
		case *dns.MX:
			x.Mx = strings.ToLower(x.Mx)
		case *dns.NAPTR:
			x.Replacement = strings.ToLower(x.Replacement)
		case *dns.KX:
			x.Exchanger = strings.ToLower(x.Exchanger)
		case *dns.SRV:
			x.Target = strings.ToLower(x.Target)
		case *dns.DNAME:
			x.Target = strings.ToLower(x.Target)
		}
		// 6.2. Canonical RR Form. (5) - origTTL
		wire := make([]byte, dns.Len(r1)+1) // +1 to be safe(r)
		off, err1 := dns.PackRR(r1, wire, 0, nil, false)
		if err1 != nil {
			return nil, err1
		}
		wire = wire[:off]
		wires[i] = wire
	}
	sort.Sort(wires)
	for i, wire := range wires {
		if i > 0 && bytes.Equal(wire, wires[i-1]) {
			continue
		}
		buf = append(buf, wire...)
	}
	return buf, nil
}

func packSigWire(sw *dns.RRSIG, msg []byte) (int, error) {
	// copied from zmsg.go RRSIG packing
	off, err := packUint16(sw.TypeCovered, msg, 0)
	if err != nil {
		return off, err
	}
	off, err = packUint8(sw.Algorithm, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint8(sw.Labels, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint32(sw.OrigTtl, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint32(sw.Expiration, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint32(sw.Inception, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint16(sw.KeyTag, msg, off)
	if err != nil {
		return off, err
	}
	off, err = dns.PackDomainName(sw.SignerName, msg, off, nil, false)
	if err != nil {
		return off, err
	}
	return off, nil
}

func packUint8(i uint8, msg []byte, off int) (off1 int, err error) {
	if off+1 > len(msg) {
		return len(msg), errors.New("overflow packing uint8")
	}
	msg[off] = byte(i)
	return off + 1, nil
}

func packUint16(i uint16, msg []byte, off int) (off1 int, err error) {
	if off+2 > len(msg) {
		return len(msg), errors.New("overflow packing uint16")
	}
	binary.BigEndian.PutUint16(msg[off:], i)
	return off + 2, nil
}


func packUint32(i uint32, msg []byte, off int) (off1 int, err error) {
	if off+4 > len(msg) {
		return len(msg), errors.New("overflow packing uint32")
	}
	binary.BigEndian.PutUint32(msg[off:], i)
	return off + 4, nil
}
