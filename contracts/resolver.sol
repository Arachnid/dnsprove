contract Resolver {
    function addr(bytes32 node) external view returns(address);
    function supportsInterface(bytes4 interfaceID) external pure returns (bool);
}
