package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := (left + right) / 2
	for left != mid {
		if nums[left] > nums[mid] {
			//mid - right为有序区间
			if target >= nums[mid] && target <= nums[right] {
				if target == nums[mid] {
					return mid
				}
				if target == nums[right] {
					return right
				}
				left = mid
			} else {
				right = mid
			}
		} else {
			//left - mid为有序区间
			if target >= nums[left] && target <= nums[mid] {
				if target == nums[left] {
					return left
				}
				if target == nums[mid] {
					return mid
				}
				right = mid
			} else {
				left = mid
			}
		}
		mid = (left + right) / 2
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return -1
}
