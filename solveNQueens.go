package main

func solveNQueens(n int) [][]string {
	queen := make([][2]int, 0)
	anslist := [][]string{}
	isRightPlace := func(x int, y int) bool {
		for _, pos := range queen {
			if y == pos[1] || x-pos[0] == y-pos[1] || x-pos[0] == pos[1]-y {
				return false
			}
		}
		return true
	}
	var backtraking func(int)
	backtraking = func(index int) {
		if index == n {
			ans := []string{}
			for _, pos := range queen {
				str := ""
				for i := 0; i < n; i++ {
					if i == pos[1] {
						str = str + "Q"
					} else {
						str = str + "."
					}
				}
				ans = append(ans, str)
			}
			anslist = append(anslist, ans)
			return
		}
		for j := 0; j < n; j++ {
			if isRightPlace(index, j) {
				queen = append(queen, [2]int{index, j})
				backtraking(index + 1)
				queen = queen[:len(queen)-1]
			}
		}
	}
	backtraking(0)
	return anslist
}
