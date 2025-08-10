// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// 引入 OpenZeppelin ERC721 库
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

// 创建 MyNFT 合约
contract WenNFT is ERC721URIStorage, Ownable {
    uint256 private _tokenIds; // 用于记录NFT的ID

    // 构造函数设置合约名称和符号
    constructor() ERC721("MyNFT", "MNFT") {}

    // mintNFT 函数：铸造NFT并关联元数据URI
    function mintNFT(
        address recipient,
        string memory tokenURI
    ) public onlyOwner returns (uint256) {
        _tokenIds++; // 增加token ID
        uint256 newItemId = _tokenIds; // 设置新NFT的ID

        _mint(recipient, newItemId); // 铸造NFT
        _setTokenURI(newItemId, tokenURI); // 设置NFT的元数据

        return newItemId; // 返回NFT ID
    }
}
