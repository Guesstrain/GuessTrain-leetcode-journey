package main

import "strconv"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	ans := ""

	for i := len(num2) - 1; i >= 0; i-- {
		curr := ""
		for j := len(num2) - 1; j > i; j-- {
			curr += "0"
		}
		curnum2, add := int(num2[i]-'0'), 0
		for j := len(num1) - 1; j >= 0; j-- {
			curnum1 := int(num1[j] - '0')
			sum := curnum1*curnum2 + add
			curr = strconv.Itoa(sum%10) + curr
			add = sum / 10
		}
		for add != 0 {
			curr = strconv.Itoa(add%10) + curr
			add = add / 10
		}
		ans = stringsum(ans, curr)
	}
	return ans
}

func stringsum(s1 string, s2 string) string {
	i, j := len(s1)-1, len(s2)-1
	add, ans := 0, ""
	for ; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(s1[i] - '0')
		}
		if j >= 0 {
			y = int(s2[j] - '0')
		}
		sum := x + y + add
		ans = strconv.Itoa(sum%10) + ans
		add = sum / 10
	}
	return ans
}

func multiplyV2(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	ansArray := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			ansArray[i+j+1] += x * y
		}
	}

	for i := m + n - 1; i > 0; i-- {
		ansArray[i-1] += ansArray[i] / 10
		ansArray[i] %= 10
	}

	ans := ""
	idx := 0
	if ansArray[0] == 0 {
		idx = 1
	}

	for ; idx < m+n; idx++ {
		ans += strconv.Itoa(ansArray[idx])
	}

	return ans
}
