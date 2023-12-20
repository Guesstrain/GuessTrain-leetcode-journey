package main

func strStr(haystack string, neddle string) int {
	if neddle == "" {
		return 0
	}
	n := len(neddle)
	lps := make([]int, n)

	prevLPS, i := 0, 1
	for i < n {
		if neddle[i] == neddle[prevLPS] {
			lps[i] = prevLPS + 1
			i++
			prevLPS++
		} else if prevLPS == 0 {
			lps[i] = 0
			i++
		} else {
			prevLPS = lps[prevLPS-1]
		}
	}

	i, j := 0, 0
	for i < len(haystack) {
		if haystack[i] == neddle[j] {
			i++
			j++
		} else if j == 0 {
			i++
		} else {
			j = lps[j-1]
		}
		if j == len(neddle) {
			return i - j
		}
	}
	return -1
}
