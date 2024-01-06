package main

func nextPermutation(nums []int) []int {
	n := len(nums)
	i := n - 1
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}
	if i == 0 {
		for j := 0; j < n/2; j++ {
			temp := nums[j]
			nums[j] = nums[n-j-1]
			nums[n-j-1] = temp
		}
		return nums
	} else {
		m := i - 1
		for i < n-1 && nums[i+1] > nums[m] {
			i++
		}
		// if i == n-1 {
		// 	i++
		// }
		temp := nums[m]
		nums[m] = nums[i]
		nums[i] = temp
		l := n - m - 1
		for j := m + 1; j < l/2+m+1; j++ {
			temp = nums[j]
			nums[j] = nums[n+m-j]
			nums[n+m-j] = temp
		}
		return nums
	}
}
