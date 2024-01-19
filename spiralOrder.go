package main

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	anslist := make([]int, 0)
	for i := 0; i < (m+1)/2 && i < (n+1)/2; i++ {
		for k := i; k < n-i; k++ {
			anslist = append(anslist, matrix[i][k])
		}
		for k := i + 1; k < m-i; k++ {
			anslist = append(anslist, matrix[k][n-1-i])
		}
		if m-1-i > i {
			for k := n - 2 - i; k >= i; k-- {
				anslist = append(anslist, matrix[m-1-i][k])
			}
			if n-1-i > i {
				for k := m - 2 - i; k > i; k-- {
					anslist = append(anslist, matrix[k][i])
				}
			}
		}
	}
	return anslist
}
