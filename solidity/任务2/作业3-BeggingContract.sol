// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract BeggingContract {
    address public owner;
    mapping(address => uint256) public donations;

    constructor() {
        owner = msg.sender;
    }

    //接收捐赠的函数
    function donate() public payable {
        require(msg.value > 0, "Donation amount must be greater than 0");
        donations[msg.sender] += msg.value;
    }
     //提取捐赠的函数
    function withdraw() public {
        require(msg.sender == owner, "Only the owner can withdraw funds");
        payable(owner).transfer(address(this).balance);  //？什么作用
    }

    //查询某个地址的捐赠金额
    function getDonation(address donor) public view returns (uint256) {
        return donations[donor];
    }

    //fallback 和 receive 防止意外发送EHT的失败
    receive() external payable {
        donate();
    }
     
    fallback() external payable {
        donate();
    }

    

}
