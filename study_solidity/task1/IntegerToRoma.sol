// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
/*
用 solidity 实现罗马数字转数整数
题目描述在 https://leetcode.cn/problems/integer-to-roman/description/
*/
contract IntegerToRoman {
    // 主函数：整数转罗马数字
    function intToRoman(uint256 num) public pure returns (string memory) {
        require(num > 0 && num <= 3999, "Out of range (1-3999)");

        // 罗马数字表
        string[13] memory romans = [
            "M",
            "CM",
            "D",
            "CD",
            "C",
            "XC",
            "L",
            "XL",
            "X",
            "IX",
            "V",
            "IV",
            "I"
        ];

        // 对应的阿拉伯数字值
        uint256[13] memory values = [
            uint256(1000),
            uint256(900),
            uint256(500),
            uint256(400),
            uint256(100),
            uint256(90),
            uint256(50),
            uint256(40),
            uint256(10),
            uint256(9),
            uint256(5),
            uint256(4),
            uint256(1)
        ];

        // 动态拼接结果
        bytes memory result;

        for (uint256 i = 0; i < values.length; i++) {
            while (num >= values[i]) {
                result = abi.encodePacked(result, romans[i]);
                num -= values[i];
            }
        }

        return string(result);
    }
}
