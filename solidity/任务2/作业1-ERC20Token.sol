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
    
    // 转账事件
    event Transfer(address indexed from, address indexed to, uint256 value);
    // 授权事件
    event Approval(address indexed owner, address indexed spender, uint256 value);
    
    // 构造函数：初始化代币信息并给部署者 mint 初始代币
    constructor(
        string memory _name,
        string memory _symbol,
        uint256 _initialSupply
    ) {
        
        name = _name;
        symbol = _symbol;
        owner = msg.sender;
        // 初始代币分配给合约部署者
        mint(msg.sender, _initialSupply * (10 **uint256(decimals)));
    }
    
    // 权限控制修饰符：仅所有者可执行
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }
    
    // 转账功能
    function transfer(address to, uint256 value) public returns (bool success) {
        require(balanceOf[msg.sender] >= value, "Insufficient balance");
        require(to != address(0), "Invalid recipient address");
        
        balanceOf[msg.sender] -= value;
        balanceOf[to] += value;
        
        emit Transfer(msg.sender, to, value);
        return true;
    }
    
    // 授权功能
    function approve(address spender, uint256 value) public returns (bool success) {
        require(spender != address(0), "Invalid spender address");
        
        allowance[msg.sender][spender] = value;
        
        emit Approval(msg.sender, spender, value);
        return true;
    }
    
    // 代扣转账功能
    function transferFrom(
        address from,
        address to,
        uint256 value
    ) public returns (bool success) {
        require(balanceOf[from] >= value, "Insufficient balance");
        require(allowance[from][msg.sender] >= value, "Allowance exceeded");
        require(to != address(0), "Invalid recipient address");
        
        balanceOf[from] -= value;
        balanceOf[to] += value;
        allowance[from][msg.sender] -= value;
        
        emit Transfer(from, to, value);
        return true;
    }
    
    // // 增发代币功能（仅所有者）
    function mint(address to, uint256 amount) public onlyOwner {
        require(to != address(0), "Invalid address");
        require(amount > 0, "Amount must be greater than 0");
        
        totalSupply += amount;
        balanceOf[to] += amount;
        
        emit Transfer(address(0), to, amount);
    }
}
