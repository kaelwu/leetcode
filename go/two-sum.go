package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 26))
}

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for k, v := range numbers {
		m[v] = k + 1
	}
	for k1, v := range numbers {
		if k2, ok := m[target-v]; ok {
			return []int{k1 + 1, k2}
		}
	}
	return []int{}
}
