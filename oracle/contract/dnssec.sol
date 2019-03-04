contract DNSSEC {
    event RRSetUpdated(bytes name, bytes rrset);

    bytes public anchors;

    function rrdata(uint16 dnstype, bytes calldata name) external view returns(uint32 inception, uint64 inserted, bytes20 hash);
    function submitRRSet(bytes calldata input, bytes calldata sig, bytes calldata proof) external;
    function submitRRSets(bytes calldata data, bytes calldata _proof) external returns (bytes memory);
    function deleteRRSet(uint16 deletetype, bytes calldata deletename, bytes calldata nsec, bytes calldata sig, bytes calldata proof) external;
}
