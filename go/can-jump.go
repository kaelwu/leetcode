package main

func canJump(nums []int) bool {
	canReach := 0
	for i := 0; i < len(nums); i++ {
		if i > canReach || canReach > len(nums)-1 {
			break
		}
		canReach = max(canReach, i+nums[i])
	}
	return canReach >= len(nums)-1
}
