pragma solidity^0.6.1;

contract calldemo {
    uint256  count;
    constructor() public {
        count = 2020;
    }
    function setCount(uint256 _count) external {
        count = _count;
    }
    function getCount() public view returns(uint256) {
        return count;
    }
}