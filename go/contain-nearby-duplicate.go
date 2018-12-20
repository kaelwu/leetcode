package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

func containsDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for k1, v := range nums {
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = k1
		}
	}
	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for k1, v := range nums {
		if k2, ok := m[v]; ok {
			if k1-k2 <= k {
				return true
			}
		}
		m[v] = k1
	}
	return false
}
