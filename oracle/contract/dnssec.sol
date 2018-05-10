contract DNSSEC {
  event RRSetUpdated(bytes indexed nameHash, bytes name);

  function rrset(uint16 class, uint16 dnstype, bytes name) public constant returns(uint32 inception, uint64 inserted, bytes rrs);
  function submitRRSet(uint16 class, bytes name, bytes input, bytes sig) public;
}
