package main

import "fmt"

func main() {
	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))
}

//func moveZeroes(nums []int) []int {
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == 0 {
//			for j := i; j < len(nums); j++ {
//				if nums[j] != 0 {
//					nums[i], nums[j] = nums[j], nums[i]
//					break
//				}
//			}
//		}
//	}
//
//	return nums
//}

func moveZeroes(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
	return nums
}
