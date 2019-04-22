package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-3]+nums[i])
	}
	max := 0
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func robTwo(nums []int) int {

}
