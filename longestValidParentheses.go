package main

func longestValidParentheses(s string) int {
	start := []int{}
	anslist := make([]int, 0)
	rstart := 0
	for i, ch := range s {
		if string(ch) == "(" {
			start = append(start, i)
		} else if string(ch) == ")" {
			lstart := len(start)
			if lstart == 0 {
				rstart = i
			} else if lstart != 0 {
				start = start[:lstart-1]
				if len(start) == 0 {
					l := i - rstart
					anslist = append(anslist, l)
				} else {
					l := i - start[len(start)-1]
					anslist = append(anslist, l)
				}
			}
		}
	}

	ansmax := 0
	for _, ans := range anslist {
		if ans > ansmax {
			ansmax = ans
		}
	}

	return ansmax
}

func longestValidParenthesesv2(s string) int {
	m := len(s)
	dp := make([]int, m+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		if string(s[i-1]) == "(" {
			dp[i] = 0
		} else {
			if i == 1 {
				dp[i] = 0
				continue
			}
			if string(s[i-2]) == "(" {
				dp[i] = dp[i-2] + 2
			} else {
				if i-dp[i-1]-2 >= 0 && string(s[i-dp[i-1]-2]) == "(" {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				} else {
					dp[i] = 0
				}
			}
		}
	}
	max := 0
	for _, d := range dp {
		if d > max {
			max = d
		}
	}
	return max
}

func longestValidParenthesesV3(s string) int {
	m := len(s)
	flag, start, maxans := 0, m-1, 0
	for i := m - 1; i >= 0; i-- {
		if string(s[i]) == "(" {
			flag--
			if flag == 0 {
				l := start - i + 1
				if l > maxans {
					maxans = l
				}
			}
		} else {
			if flag < 0 {
				flag = 0
				start = i
			}
			flag++
		}
	}

	flag, start = 0, 0
	for i := 0; i < m; i++ {
		if string(s[i]) == "(" {
			if flag < 0 {
				flag = 0
				start = i
			}
			flag++
		} else {
			flag--
			if flag == 0 {
				l := i - start + 1
				if l > maxans {
					maxans = l
				}
			}
		}
	}
	return maxans
}
