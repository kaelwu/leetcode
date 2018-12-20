package main

import "fmt"

func main() {
	fmt.Println(10 ^ 3)
}

func twoSum(nums []int, target int) []int {
	b := make(map[int]int)
	res := make([]int, 0)
	for k1, v1 := range nums {
		b[target-v1] = k1
	}
	fmt.Println(b)
	for k2, v2 := range nums {
		if _, ok := b[v2]; ok {
			if b[v2] == k2 {
				continue
			} else {
				res = append(res, b[v2], k2)
				return res
			}
		}
	}
	return res
}
