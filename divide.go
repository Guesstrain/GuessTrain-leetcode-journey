package main

import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	positive := 1
	if dividend*divisor < 0 {
		positive = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	ans, step := 0, 0
	for dividend >= divisor {
		if dividend > divisor<<step {
			dividend -= divisor << step
			ans += 1 << step
			step++
		} else {
			step--
		}
	}
	return ans * positive
}
