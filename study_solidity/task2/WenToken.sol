// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract WenToken {
    string public name = "WenToken";
    string public symbol = "WTK";
    uint8 public decimals = 18;
    uint256 public totalSupply;

    address public owner;

    mapping(address => uint256) private balances;
    mapping(address => mapping(address => uint256)) private allowances;

    // 事件定义
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );

    constructor() {
        owner = msg.sender;
    }

    // 查询余额
    function balanceOf(address account) public view returns (uint256) {
        return balances[account];
    }

    // 转账
    function transfer(address to, uint256 amount) public returns (bool) {
        require(balances[msg.sender] >= amount, "Insufficient balance");
        _transfer(msg.sender, to, amount);
        return true;
    }

    // 授权 spender 从 msg.sender 花费金额
    function approve(address spender, uint256 amount) public returns (bool) {
        allowances[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    // 被授权者调用，转账代扣
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) public returns (bool) {
        require(allowances[from][msg.sender] >= amount, "Not allowed");
        require(balances[from] >= amount, "Insufficient balance");

        allowances[from][msg.sender] -= amount;
        _transfer(from, to, amount);
        return true;
    }

    // 授权额度查询
    function allowance(
        address _owner,
        address spender
    ) public view returns (uint256) {
        return allowances[_owner][spender];
    }

    // 增发（只有合约部署者能调用）
    function mint(address to, uint256 amount) public {
        require(msg.sender == owner, "Only owner can mint");
        balances[to] += amount;
        totalSupply += amount;
        emit Transfer(address(0), to, amount);
    }

    // 内部转账函数
    function _transfer(address from, address to, uint256 amount) internal {
        require(to != address(0), "Cannot transfer to zero address");

        balances[from] -= amount;
        balances[to] += amount;
        emit Transfer(from, to, amount);
    }
}
