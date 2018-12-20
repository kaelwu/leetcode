package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func threeSum(nums []int) [][]int {
	m := make(map[int]int)
	i := 0
	for _, v := range nums {
		m[i] = 0 - v
		i++
	}
	for k, v := range m {
		for i := 0; i < len(nums); i++ {
			if k == v {
				continue
			}
		}
	}
}
