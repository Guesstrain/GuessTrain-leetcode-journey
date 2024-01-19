package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to GuessTrain's leetcode journey")

	//17.电话号码的字母组合
	//testLetter := "369"
	//fmt.Println(letterCombinations(testLetter))

	//18.四数之和
	// testnums := []int{2, 2, 2, 2, 2}
	// fmt.Println(fourSum(testnums, 8))

	// 19.删除链表倒数第N个节点
	// testListNode := ListNode{1, nil}
	// testListNode.Next = &ListNode{2, nil}
	// fmt.Println(removeNthFromEnd(&testListNode, 1))

	// 22.生成括号
	// testnum := 3
	// fmt.Println(generateParenthesis((testnum)))

	// 28.找出字符串中第一个匹配项下标
	// haystack := "ababcaababcaabc"
	// neddle := "ababcaabc"
	// fmt.Println(strStr(haystack, neddle))

	// 29.两数相除
	// dividend, divisor := 10, 3
	// fmt.Println(divide(dividend, divisor))

	// 30. 串联所有单词的子串
	// s := "a"
	// words := []string{"a"}
	// fmt.Println(findSubstring(s, words))

	// 50. 字母异位词分组
	// strs := []string{"", "", ""}
	// fmt.Println(groupAnagrams(strs))

	// 53. 最大字数和
	// nums := []int{100, 101, 102, -1, -88, -22, 200}
	// fmt.Println(maxSubArray(nums))

	//31. 下一个排列
	// nums := []int{1, 5, 1}
	// fmt.Println(nextPermutation(nums))

	//32. 最长有效括号
	// s := ")()())"
	// fmt.Println(longestValidParenthesesV3(s))

	//33. 搜索旋转排序数组
	// nums := []int{1, 3}
	// fmt.Println(search(nums, 3))

	//34. 在排序数组中查找元素的第一个和最后一个位置
	// nums := []int{1, 2, 3}
	// fmt.Println(searchRange(nums, 2))

	//38. 外观数列
	// n := 4
	// fmt.Println(countAndSayV2(n))

	//39. 组合总数
	// c := []int{8, 7, 4, 3}
	// fmt.Println(combinationSumV2(c, 11))

	//40. 组合总数2
	// c := []int{10, 1, 2, 7, 6, 1, 5}
	// fmt.Println(combinationSum2(c, 8))

	//41. 缺失的第一个整数
	// c := []int{3, 4, -1, 1}
	// fmt.Println(firstMissingPositiveV2(c))

	//42. 接雨水
	// c := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// fmt.Println(trapV3(c))

	//43. 字符串相乘
	// num1, num2 := "123", "456"
	// fmt.Println(multiplyV2(num1, num2))

	//46. 全排列
	// nums := []int{1, 2, 3}
	// fmt.Println(permute(nums))

	//47. 全排列2
	// nums := []int{1, 1, 2}
	// fmt.Println(permuteUnique(nums))

	//48. 旋转图像
	// matrix := make([][]int, 0)
	// rotate(matrix)

	//51. N皇后
	// fmt.Println(solveNQueens(4))

	//54. 螺旋矩阵
	// matrix := [][]int{[]int{2, 5, 8}, []int{4, 0, -1}}
	// matrix := [][]int{[]int{7, 9, 6}}
	// fmt.Println(spiralOrder(matrix))

	//57. 插入区间
	intervals := [][]int{[]int{1, 3}, []int{6, 9}}
	newInterval := []int{2, 5}
	fmt.Println(insert2(intervals, newInterval))

}
