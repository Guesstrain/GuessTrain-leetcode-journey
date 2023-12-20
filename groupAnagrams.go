package main

func groupAnagrams(strs []string) [][]string {
	ans := make([][]string, 0)
	for i := 0; i < len(strs); i++ {
		sameCharLists := make([]string, 0)
		for j := i + 1; j < len(strs); j++ {
			if isSameCharacter(strs[i], strs[j]) {
				sameCharLists = append(sameCharLists, strs[j])
				strs = append(strs[:j], strs[j+1:]...)
				j--
			}
		}
		sameCharLists = append(sameCharLists, strs[i])
		ans = append(ans, sameCharLists)
	}
	return ans
}

func isSameCharacter(str1 string, str2 string) bool {
	differ := map[rune]int{}
	for _, char := range str1 {
		differ[char]++
	}
	for _, char := range str2 {
		differ[char]--
		if differ[char] == 0 {
			delete(differ, char)
		}
	}
	if len(differ) == 0 {
		return true
	}
	return false
}
