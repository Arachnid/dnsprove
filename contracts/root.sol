contract Root {
    function oracle() public view returns(address);
    function proveAndRegisterTLD(bytes calldata name, bytes calldata input, bytes calldata proof) external;
    function proveAndRegisterDefaultTLD(bytes calldata name, bytes calldata input, bytes calldata proof) external;
    function registerTLD(bytes memory name, bytes memory proof) public;
    function supportsInterface(bytes4 interfaceID) external pure returns (bool);
}
