contract DNSSEC {
    event RRSetUpdated(bytes name, bytes rrset);

    bytes public anchors;

    function rrdata(uint16 dnstype, bytes name) public constant returns(uint32 inception, uint64 inserted, bytes20 hash);
    function submitRRSet(bytes input, bytes sig, bytes proof) public;
    function deleteRRSet(uint16 deletetype, bytes deletename, bytes nsec, bytes sig, bytes proof) public;
}
