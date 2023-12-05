package main

func letterCombinations(digits string) []string {
	dic := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz"}
	ans := make([]string, 0)

	for _, numstr := range digits {
		if len(ans) == 0 {
			for _, char := range dic[string(numstr)] {
				ans = append(ans, string(char))
			}
		} else {
			temp := make([]string, 0)
			for _, char1 := range ans {
				for _, char2 := range dic[string(numstr)] {
					temp = append(temp, char1+string(char2))
				}
			}
			ans = temp
		}
	}
	return ans
}
