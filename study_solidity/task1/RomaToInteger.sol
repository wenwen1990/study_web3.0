// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
/*
用 solidity 实现整数转罗马数字
题目描述在 https://leetcode.cn/problems/roman-to-integer/description/3.
*/
contract RomaToInteger {
    mapping(bytes1 => uint256) internal map;

    constructor() {
        map["I"] = 1;
        map["V"] = 5;
        map["X"] = 10;
        map["L"] = 50;
        map["C"] = 100;
        map["D"] = 500;
        map["M"] = 1000;
    }

    function cal(string memory str) public view returns (uint256 num) {
        bytes memory strBytes = bytes(str);
        for (uint256 i = 0; i < strBytes.length; i++) {
            uint256 current = map[strBytes[i]];
            uint256 next = (i + 1 < strBytes.length) ? map[strBytes[i + 1]] : 0;

            if (current < next) {
                num += (next - current);
                i++; // 跳过下一个字符，因为它已经被处理过
            } else {
                num += current;
            }
        }
    }
}
