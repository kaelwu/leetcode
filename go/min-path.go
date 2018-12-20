package main

import "fmt"

func main() {
	m := [][]int{}
	m = append(m, []int{1, 3, 1})
	m = append(m, []int{1, 5, 1})
	m = append(m, []int{4, 2, 1})
	fmt.Println(minPathSum(m))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		tmp := make([]int, m)
		d[i] = tmp
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				d[i][j] = grid[i][j]
			} else if i == 0 && j != 0 {
				d[i][j] = d[i][j-1] + grid[i][j]
			} else if j == 0 && i != 0 {
				d[i][j] = d[i-1][j] + grid[i][j]
			} else {
				d[i][j] = min(d[i-1][j], d[i][j-1]) + grid[i][j]
			}
		}
	}
	return d[n-1][m-1]
}
