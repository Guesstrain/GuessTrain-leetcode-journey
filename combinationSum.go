package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	anslist := make([][]int, 0)

	if target < candidates[0] {
		return [][]int{}
	}

	index := sort.SearchInts(candidates, target)
	if index < len(candidates) && candidates[index] == target {
		anslist = append(anslist, []int{candidates[index]})
	}
	for i := index - 1; i >= 0; i-- {
		newtarget := target - candidates[i]
		ansnt := combinationSum(candidates[:i+1], newtarget)
		for _, ans := range ansnt {
			anscur := append(ans, candidates[i])
			anslist = append(anslist, anscur)
		}
	}
	return anslist
}

func combinationSumV2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	comb := []int{}
	var dfs func(int, int)
	dfs = func(target int, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == candidates[idx] {
			comb = append(comb, candidates[idx])
			ans = append(ans, append([]int(nil), comb...))
			comb = comb[:len(comb)-1]
		}
		dfs(target, idx+1)
		if target > candidates[idx] {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return ans
}
