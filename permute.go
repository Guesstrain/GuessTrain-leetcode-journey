package main

func permute(nums []int) [][]int {
	anslist := make([][]int, 0)
	showmap := map[int]int{}
	var backtracking func(int)
	backtracking = func(index int) {
		if index == len(nums) {
			ans := make([]int, index)
			for k, v := range showmap {
				ans[v] = k
			}
			anslist = append(anslist, ans)
		}
		for i := 0; i < len(nums) && index < len(nums); i++ {
			_, ok := showmap[nums[i]]
			if !ok {
				showmap[nums[i]] = index
				backtracking(index + 1)
				delete(showmap, nums[i])
			}
		}
	}
	backtracking(0)
	return anslist
}
