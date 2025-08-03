// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

/*
 合并两个有序数组 (Merge Sorted Array)
题目描述：将两个有序数组合并为一个有序数组。
*/
contract MergetSortedArray {
    function mergeArr(
        uint[] memory arr1,
        uint[] memory arr2
    ) public pure returns (uint[] memory res) {
        uint l1 = arr1.length;
        uint l2 = arr2.length;
        res = new uint[](l1 + l2);
        for (uint i = 0; i < l1; i++) {
            res[i] = arr1[i];
        }
        for (uint i = 0; i < l2; i++) {
            res[l1 + i] = arr2[i];
        }
        res = bubbleSort(res);
    }

    function bubbleSort(uint[] memory arr) public pure returns (uint[] memory) {
        for (uint i = 0; i < arr.length; i++) {
            for (uint256 j = 0; j < arr.length - 1 - i; j++) {
                if (arr[j] > arr[j + 1]) {
                    uint temp = arr[j + 1];
                    arr[j + 1] = arr[j];
                    arr[j] = temp;
                }
            }
        }
        return arr;
    }
}
