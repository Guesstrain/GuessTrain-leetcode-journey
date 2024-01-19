package main

import "sort"

func permuteUnique(nums []int) [][]int {
	m, anslist := len(nums), make([][]int, 0)
	choosenum := nums
	sort.Ints(choosenum)
	var backtracking func(int)
	backtracking = func(index int) {
		if index == m {
			anslist = append(anslist, append([]int{}, choosenum...))
		}
		for i := index; i < m; i++ {
			if i > index && choosenum[i] == choosenum[i-1] {
				continue
			}
			if index != i {
				temp := choosenum[i]
				for j := i; j > index; j-- {
					choosenum[j] = choosenum[j-1]
				}
				choosenum[index] = temp
				backtracking(index + 1)
				temp = choosenum[index]
				for j := index; j < i; j++ {
					choosenum[j] = choosenum[j+1]
				}
				choosenum[i] = temp
			} else {
				backtracking(index + 1)
			}
		}
	}
	backtracking(0)
	return anslist
}
