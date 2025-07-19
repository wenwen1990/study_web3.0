package main

import (
	"fmt"
	"sort"
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

	//fmt.Println(isPalindrome(121))
	//fmt.Println(isPalindrome(-121))
	//fmt.Println(isPalindrome(10))

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

	//fmt.Println(plusOne([]int{1, 2, 3}))
	//fmt.Println(plusOne([]int{1, 2, 3, 9, 9, 9}))
	//fmt.Println(plusOne([]int{9, 9, 9}))

	// fmt.Println(removeDuplicates([]int{1, 1, 2}))
	// fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	// fmt.Println(merge([][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}))
	// fmt.Println(merge([][]int{[]int{1, 4}, []int{4, 5}}))

	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
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

/*
26. 删除有序数组中的重复项
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。
*/
func removeDuplicates(nums []int) int {
	//解题思路重点在于双指针快慢指针过滤数组数据，记录慢指针
	length := len(nums)
	if length == 0 {
		return 0
	}
	var slow = 0
	for fast := 1; fast < length; fast++ {
		slowValue := nums[slow]
		fastValue := nums[fast]
		if slowValue != fastValue {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

/*
56. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/
func merge(intervals [][]int) [][]int {
	var res [][]int
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for idx := 0; idx < len(intervals); idx++ {
		targetArr := intervals[idx]
		if len(res) == 0 || res[len(res)-1][1] < targetArr[0] {
			res = append(res, targetArr)
			continue
		}
		res[len(res)-1][1] = max(res[len(res)-1][1], targetArr[1])
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
1. 两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
你可以按任意顺序返回答案。
*/
func twoSum(nums []int, target int) []int {
	var res = []int{}
	var value2idxMap = make(map[int]int)
	for idx, value := range nums {
		targetValue := nums[idx]
		otherValue := target - targetValue
		otherIdx, exist := value2idxMap[otherValue]
		if exist {
			res = append(res, idx)
			res = append(res, otherIdx)
			return res
		}
		value2idxMap[value] = idx
	}
	return res
}
