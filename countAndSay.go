package main

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	count := make([][]int, 0)
	for _, v := range countAndSay(n - 1) {
		num := int(v - '0')
		if len(count) == 0 {
			count = append(count, []int{num, 1})
		} else {
			if num == count[len(count)-1][0] {
				count[len(count)-1][1]++
			} else {
				count = append(count, []int{num, 1})
			}
		}
	}
	ans := ""
	for _, row := range count {
		i, j := row[0], row[1]
		ans = ans + strconv.Itoa(j*10+i)
	}
	return ans
}

func countAndSayV2(n int) string {
	prev := "1"

	for i := 2; i <= n; i++ {
		var builder strings.Builder
		for j, start := 0, 0; j < len(prev); start = j {
			for j < len(prev) && prev[j] == prev[start] {
				j++
			}
			builder.WriteString(strconv.Itoa(j - start))
			builder.WriteByte(prev[start])
		}
		prev = builder.String()
	}
	return prev
}
