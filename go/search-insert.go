package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3}, 4))
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		} else {
			return 1
		}
	}
	mid := int(len(nums) / 2)
	if nums[mid] > target {
		return searchInsert(nums[0:mid], target)
	} else if nums[mid] < target {
		return mid + searchInsert(nums[mid+1:], target) + 1
	} else {
		return mid
	}
}
