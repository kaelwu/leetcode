package main

import "fmt"

func main() {
	fmt.Println(rotate([]int{1, 2, 3}, 2))
}

func rotate(nums []int, k int) []int {
	for i := 1; i <= k; i++ {
		old := nums[0]
		for c := range nums {
			if c+1 >= len(nums) {
				c = -1
			}
			nums[c+1], old = old, nums[c+1]
		}
	}
	return nums
}
