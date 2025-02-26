// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract Token is ERC20 {
    address public owner;
    constructor() ERC20("R-Token", "RTN") {
        owner = msg.sender;
    }

    modifier onlyOwner(){
        require(msg.sender == owner, "not contract owner");
        _;
    }

    function mint(uint value, address to) public onlyOwner {
        _mint(to, value);
    }
}