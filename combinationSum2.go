package main

func combinationSum2(candidates []int, target int) [][]int {
	comb := make([]int, 0)
	anslist := make([][]int, 0)
	var dfs func(int, int)
	dfs = func(target int, idx int) {
		if target == 0 {
			anslist = append(anslist, append([]int(nil), comb...))
			return
		}

		if idx == len(candidates) {
			return
		}

		dfs(target, idx+1)

		if target > 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx+1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return anslist
}
