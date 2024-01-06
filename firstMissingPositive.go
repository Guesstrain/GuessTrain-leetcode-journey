package main

func firstMissingPositive(nums []int) int {
	m := len(nums)
	for i, num := range nums {
		if num <= 0 {
			nums[i] = m + 1
		}
	}

	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if num < m+1 && num > 0 {
			if nums[num-1] > 0 {
				nums[num-1] = -nums[num-1]
			}
		}
	}

	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}
	return m + 1
}

func firstMissingPositiveV2(nums []int) int {
	n := len(nums)
	for i, _ := range nums {
		for nums[i] > 0 && nums[i] < n+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}

	return n + 1
}
