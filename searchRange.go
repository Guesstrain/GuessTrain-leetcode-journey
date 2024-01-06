package main

import "sort"

func searchRange(nums []int, target int) []int {
	leftindex := binarySearch(nums, target, true)
	rightindex := binarySearch(nums, target, false)
	if leftindex == len(nums) || nums[leftindex] != target {
		return []int{-1, -1}
	} else {
		return []int{leftindex, rightindex - 1}
	}
}

func binarySearch(nums []int, target int, lower bool) int {
	left, right, ans := 0, len(nums)-1, len(nums)
	mid := (left + right) / 2
	for left <= right {
		if nums[mid] > target || (lower && nums[mid] >= target) {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return ans
}

func searchRangeV2(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}
