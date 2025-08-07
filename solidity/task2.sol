作业 1：ERC20 代币
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



作业2：在测试网上发行一个图文并茂的 NFT
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts@4.9.0/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts@4.9.0/access/Ownable.sol";
contract MyNFT is ERC721URIStorage, Ownable {
    uint256 private _tokenIds;

    constructor() ERC721("MyNFT", "SamNFT") {}

    function mintNFT(address recipient, string memory tokenURI)
        public onlyOwner
        returns (uint256)
    {
        _tokenIds++;
        uint256 newItemId = _tokenIds;

        _mint(recipient, newItemId);
        _setTokenURI(newItemId, tokenURI);

        return newItemId;
    }
}



作业3：编写一个讨饭合约
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


