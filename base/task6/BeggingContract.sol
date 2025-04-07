// SPDX-License-Identifier: MIT

pragma solidity ^0.8;

import "@openzeppelin/contracts/access/Ownable.sol";


contract BeggingContract is Ownable {
    // 捐赠金额
    mapping(address => uint) private  donateMap;
    // 捐赠事件
    event Donation(address indexed donor, uint256 amount);
    // 捐赠开始时间
    uint256 public startTime;
    // 捐赠结束时间
    uint256 public endTime;
    // 排序

    constructor(uint256 _startTime, uint256 _endTime) Ownable(msg.sender) {
        startTime = _startTime;
        endTime = _endTime;
    }


    function getDonation(address account) external view  returns (uint256) {
        require(account != address(0), "account is zero");
        return donateMap[account];
    }


    function donate() public payable   {
        require(block.timestamp >= startTime && block.timestamp <= endTime, "Donation is not allowed at this time");
        require(msg.value > 0, "Donation amount must be greater than 0");

        donateMap[msg.sender] += msg.value;
        emit Donation(msg.sender, msg.value);
    }

    // 允许合约所有者提取所有资金
    function withdraw() external onlyOwner {
        uint256 balance = address(this).balance;
        payable(owner()).transfer(balance);
    }

}