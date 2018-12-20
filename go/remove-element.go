package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}

func removeElement(nums []int, val int) int {
	var i int
	for k, v := range nums {
		if v != val {
			nums[i] = nums[k]
			i++
		}
	}
	nums = nums[:i]
	return i
}
