package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	d := make([]int, len(nums))
	// 取边界
	d[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		// 公式
		d[i] = max(d[i-1]+nums[i], nums[i])
	}
	maxNum := d[0]
	for i := 0; i < len(d); i++ {
		if d[i] > maxNum {
			maxNum = d[i]
		}
	}
	return maxNum

}
