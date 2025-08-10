// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BeggingContract {
    address internal owner; // 合约所有者
    mapping(address => uint256) internal donations; // 记录每个捐赠者的金额

    // 部署者是合约所有者
    constructor() {
        owner = msg.sender;
    }

    event donateEvent(address indexed donator, uint256 amount);
    event withdrawEvent(address indexed owner, uint256 amount);

    function donate() public payable {
        require(msg.value > 0, "donation must gether than zero");
        donations[msg.sender] += msg.value;
        emit donateEvent(msg.sender, msg.value);
    }

    function withdraw() public {
        require(msg.sender == owner, "only owner can withdraw");
        uint256 amount = address(this).balance;
        require(amount > 0, "no funds to withdraw");
        payable(owner).transfer(amount);
        emit withdrawEvent(owner, amount);
    }

    function getDonation(address addr) public view returns (uint256) {
        return donations[addr];
    }

    receive() external payable {
        donate();
    }
}
