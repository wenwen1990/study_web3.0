package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//nums1 := []int{2, 2, 1}
	//nums2 := []int{4, 1, 2, 1, 2}
	//nums3 := []int{1}
	//fmt.Printf("%+v\n", singleNumber(nums1))
	//fmt.Printf("%+v\n", singleNumber(nums2))
	//fmt.Printf("%+v\n", singleNumber(nums3))
	//
	//fmt.Println(isPalindrome(121))
	//fmt.Println(isPalindrome(-121))
	//fmt.Println(isPalindrome(10))
	//
	//fmt.Println(isValid("()"))
	//fmt.Println(isValid("()[]{}"))
	//fmt.Println(isValid("(]"))
	//fmt.Println(isValid("([])"))
	//fmt.Println(isValid("([)]"))
	//fmt.Println(isValid("]"))

	//fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	//fmt.Println(longestCommonPrefix([]string{""}))
	//fmt.Println(longestCommonPrefix([]string{"flower", "flower", "flower", "flower"}))

	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{1, 2, 3, 9, 9, 9}))
	fmt.Println(plusOne([]int{9, 9, 9}))
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

/*
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
func isValid(s string) bool {
	arr := strings.Split(s, "")
	stack := []string{}
	for _, value := range arr {
		if value == "(" || value == "[" || value == "{" {
			stack = append(stack, value)
		} else if value == ")" {
			if len(stack) == 0 {
				return false
			}
			lastValue := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if lastValue != "(" {
				return false
			}
		} else if value == "]" {
			if len(stack) == 0 {
				return false
			}
			lastValue := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if lastValue != "[" {
				return false
			}
		} else if value == "}" {
			if len(stack) == 0 {
				return false
			}
			lastValue := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if lastValue != "{" {
				return false
			}
		}
	}
	return len(stack) == 0
}

/*
*
14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*/
func longestCommonPrefix(strs []string) string {
	res := ""
	length := len(strs)
	if length == 0 || strs[0] == "" {
		return res
	}
	firstStr := strs[0]
	for i := 0; i < len(firstStr); i++ {
		targetValue := firstStr[i]
		for j := 1; j < length; j++ {
			otherStr := strs[j]
			if len(otherStr)-1 < i {
				return res
			} else if otherStr[i] != targetValue {
				return res
			}
		}
		res += string(targetValue)
	}

	return res
}

/*
66. 加一
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
将大整数加 1，并返回结果的数字数组。
*/
func plusOne(digits []int) []int {
	length := len(digits)
	var carry = 1
	for i := length - 1; i >= 0; i-- {
		if carry == 0 {
			break
		}
		lastValue := digits[i]
		if lastValue != 9 {
			lastValue++
			digits[i] = lastValue
			carry = 0
		} else {
			digits[i] = 0
		}
	}
	if carry == 1 {
		var newArr = make([]int, length+1)
		copy(newArr[1:], digits)
		newArr[0] = 1
		return newArr
	}
	return digits
}
