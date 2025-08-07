// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ERC20Token {
    // 代币名称
    string public name;
    // 代币符号
    string public symbol;
    // 小数位数
    uint8 public decimals = 18;
    // 代币总供应量
    uint256 public totalSupply;
    
    // 存储账户余额
    mapping(address => uint256) public balanceOf;
    // 存储授权信息 (owner => spender => amount)
    mapping(address => mapping(address => uint256)) public allowance;
    
    // 合约所有者地址
    address public owner;

    
