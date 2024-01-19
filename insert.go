package main

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		intervals = append(intervals, newInterval)
		return intervals
	}
	newIntervals := make([][]int, 0)
	start, end := newInterval[0], newInterval[1]
	i := 0
	for ; i < len(intervals); i++ {
		if newInterval[0] <= intervals[i][0] || newInterval[0] <= intervals[i][1] {
			if newInterval[0] > intervals[i][0] {
				start = intervals[i][0]
			} else {
				start = newInterval[0]
			}
			j := i
			for j < len(intervals) && newInterval[1] > intervals[j][1] {
				j++
			}
			if j == len(intervals) {
				end = newInterval[1]
			} else if newInterval[1] >= intervals[j][0] {
				end = intervals[j][1]
			} else {
				end = newInterval[1]
				newIntervals = append(newIntervals, []int{start, end})
				newIntervals = append(newIntervals, intervals[j])
				i = j
				break
			}
			i = j
			newIntervals = append(newIntervals, []int{start, end})
			break
		}
		newIntervals = append(newIntervals, intervals[i])
	}
	i++
	for ; i < len(intervals); i++ {
		newIntervals = append(newIntervals, intervals[i])
	}
	if newInterval[0] > intervals[len(intervals)-1][1] {
		newIntervals = append(newIntervals, newInterval)
	}
	return newIntervals
}

func insert2(intervals [][]int, newInterval []int) (ans [][]int) {
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[0] > right {
			// 在插入区间的右侧且无交集
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {
			// 在插入区间的左侧且无交集
			ans = append(ans, interval)
		} else {
			// 与插入区间有交集，计算它们的并集
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !merged {
		ans = append(ans, []int{left, right})
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
