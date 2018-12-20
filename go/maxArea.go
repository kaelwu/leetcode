package main

import "fmt"

func main() {
	res := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(res)
}

func maxArea(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1
	for left < right {
		tmp := min(height[left], height[right]) * (right - left)
		if tmp > max {
			max = tmp
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func min(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

func abs(a, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}
