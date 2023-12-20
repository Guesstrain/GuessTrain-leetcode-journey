package main

var neddles []string

func findSubstring(s string, words []string) []int {
	str := ""
	neddles = []string{}
	getSubstring(words, str)
	ans := make([]int, 0)
	for _, neddle := range neddles {
		index := strStr1(s, neddle)
		for _, inde := range index {
			if !contains(ans, inde) {
				ans = append(ans, inde)
			}
		}

	}
	return ans
}

func strStr1(s string, neddle string) []int {
	anslist := make([]int, 0)
	if neddle == "" {
		return anslist
	}
	n := len(neddle)
	lps := make([]int, n)

	prevLPS, i := 0, 1
	for i < n {
		if neddle[prevLPS] == neddle[i] {
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
	for i < len(s) {
		if s[i] == neddle[j] {
			i++
			j++
		} else if j == 0 {
			i++
		} else {
			j = lps[j-1]
		}
		if j == len(neddle) {
			anslist = append(anslist, i-j)
			i = i - j + 1
			j = 0
		}
	}
	return anslist
}

func getSubstring(words []string, str string) {
	if len(words) == 0 {
		neddles = append(neddles, str)
	}

	for i, word := range words {
		newstr := str + words[i]
		newwords := make([]string, len(words)-1)
		k := 0
		for j, _ := range words {
			if j != i {
				newwords[k] = words[j]
				k++
			}
		}
		words[i] = word
		getSubstring(newwords, newstr)
	}
}

func contains(numbers []int, numberToFind int) bool {
	for _, number := range numbers {
		if number == numberToFind {
			return true
		}
	}
	return false
}
