package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	if len(nums) < 4 {
		return [][]int{}
	}

	for i := 0; i < n; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			target2 := target - (nums[i] + nums[j])
			l := n - 1
			for k := j + 1; k < n; k++ {
				if k != j+1 && nums[k] == nums[k-1] {
					continue
				}
				for k < l && nums[k]+nums[l] > target2 {
					l--
				}
				if k == l {
					break
				}
				if nums[k]+nums[l] == target2 {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[l]})
				}
			}
		}
	}
	return ans
}
