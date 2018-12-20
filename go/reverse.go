package main

import "fmt"

func main() {
	fmt.Println(2 << 30)
	nums := []int{123, -321, 120, 1563847412}
	for _, v := range nums {
		fmt.Println(reverse(v))
	}
}

func reverse(x int) int {
	nums := make([]int, 0)
	for {
		if x == 0 {
			break
		}
		tmp := x % 10
		nums = append(nums, tmp)
		x = int(x / 10)
	}
	i := 1
	sum := 0
	for j := len(nums) - 1; j >= 0; j-- {
		sum += i * nums[j]
		i = i * 10
	}
	if sum > 2<<30-1 || sum < -2<<30 {
		return 0
	}
	return sum
}
