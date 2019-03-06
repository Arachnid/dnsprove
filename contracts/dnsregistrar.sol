contract DNSRegistrar {
    function oracle() external view returns(address);
    function claim(bytes calldata name, bytes calldata proof) external;
    function proveAndClaim(bytes calldata name, bytes calldata input, bytes calldata proof) external;
    function supportsInterface(bytes4 interfaceID) external pure returns (bool);
}
