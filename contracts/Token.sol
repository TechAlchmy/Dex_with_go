// init spdx for version ^0.8.0
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

// contract Token inherits from ERC20
contract Token is ERC20{
    constructor(uint256 initialSupply) ERC20("TimCoin", "TIM") {
        _mint(msg.sender, initialSupply);
    }
}