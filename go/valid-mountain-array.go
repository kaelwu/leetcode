package main

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{9, 8, 7}))
	fmt.Println(validMountainArray([]int{6, 8, 7}))
	fmt.Println(validMountainArray([]int{3, 7, 5, 4, 0, 1, 0}))
}

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	top := 0
	for i := 0; i < len(A); i++ {
		if i+1 == len(A) {
			continue
		}
		if A[i+1] >= A[i] && top == i {
			top = i + 1
			continue
		} else if A[i+1] >= A[i] && top != i {
			return false
		} else if A[i+1] < A[i] && top == 0 {
			return false
		} else {
			continue
		}
	}
	if top == 0 || top == len(A)-1 {
		return false
	}
	return true
}
