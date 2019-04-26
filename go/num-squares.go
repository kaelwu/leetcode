package main

import "fmt"

func main() {
	fmt.Println(numSquares(12))
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		dp[i] = i
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
