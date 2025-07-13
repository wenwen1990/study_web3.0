package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums1 := []int{2, 2, 1}
	nums2 := []int{4, 1, 2, 1, 2}
	nums3 := []int{1}
	fmt.Printf("%+v\n", singleNumber(nums1))
	fmt.Printf("%+v\n", singleNumber(nums2))
	fmt.Printf("%+v\n", singleNumber(nums3))

	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}

/*
136. 只出现一次的数字
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
*/
func singleNumber(nums []int) int {
	var res int
	map1 := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		record := nums[i]
		existTimes, exist := map1[record]
		if exist {
			existTimes++
		} else {
			existTimes = 1
		}
		map1[record] = existTimes
	}

	for key, value := range map1 {
		if value == 2 {
			continue
		} else if value == 1 {
			res = key
			break
		}
	}

	return res
}

/*
9. 回文数
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*/
func isPalindrome(x int) bool {
	var res = true
	numStr := strconv.Itoa(x)
	var bytes = []byte(numStr)
	for i := 0; i < len(bytes); i++ {
		var firstStr = bytes[i]
		var lastStr = bytes[len(bytes)-1-i]
		if firstStr != lastStr {
			res = false
			break
		}
	}
	return res
}
