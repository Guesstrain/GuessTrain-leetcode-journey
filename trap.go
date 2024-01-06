package main

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	water := 0
	peak := make([]int, 0)
	k := trend(height[0], height[1])
	if k < 0 {
		peak = append(peak, 0)
	}
	for i := 2; i < len(height); i++ {
		newk := trend(height[i-1], height[i])
		if newk <= 0 && k >= 0 {
			peak = append(peak, i-1)
		}
		k = newk
	}
	if k > 0 {
		peak = append(peak, len(height)-1)
	}

	hasneg := true
	for hasneg {
		newpeak := make([]int, 0)
		leftmax, rightmax := -1, -1
		for i := 1; i < len(peak); i++ {
			k = trend(height[peak[i-1]], height[peak[i]])
			if k < 0 && leftmax == -1 {
				leftmax = i - 1
			}
			if k > 0 && leftmax >= 0 && rightmax == -1 {
				rightmax = i
			}
			if k < 0 && rightmax >= 0 {
				for j := leftmax + 1; j < rightmax; j++ {
					if height[peak[j]] < height[peak[leftmax]] && height[peak[j]] < height[peak[rightmax]] {
						peak[j] = -peak[j]
					}
				}
				leftmax = i - 1
				rightmax = -1
			}
		}
		if leftmax != -1 && rightmax != -1 {
			for j := leftmax + 1; j < rightmax; j++ {
				if height[peak[j]] < height[peak[leftmax]] && height[peak[j]] < height[peak[rightmax]] {
					peak[j] = -peak[j]
				}
			}
		}
		hasneg = false
		for i := 0; i < len(peak); i++ {
			if peak[i] >= 0 {
				newpeak = append(newpeak, peak[i])
			} else {
				hasneg = true
			}
		}
		peak = newpeak
	}

	for i := 0; i < len(peak)-1; i++ {
		peak1, peak2 := peak[i], peak[i+1]
		lower := less(height[peak1], height[peak2])
		for j := peak1 + 1; j < peak2; j++ {
			if height[j] < lower {
				water = water + lower - height[j]
			}
		}
	}
	return water
}

func trend(prev int, cur int) int {
	if prev > cur {
		return -1
	} else if prev < cur {
		return 1
	}
	return 0
}

func less(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func trapV2(height []int) int {
	n := len(height)
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0], rightMax[n-1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		if height[i] > leftMax[i-1] {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}
	for i := n - 2; i >= 0; i-- {
		if height[i] > rightMax[i+1] {
			rightMax[i] = height[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = ans + less(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}

func trapV3(height []int) (ans int) {
	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curwidth := i - left - 1
			curheight := less(h, height[left]) - height[top]
			ans += curheight * curwidth
		}
		stack = append(stack, i)
	}
	return
}
