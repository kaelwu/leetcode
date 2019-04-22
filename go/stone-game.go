package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func stoneGame(piles []int) bool {
	low := 0
	high := len(piles) - 1
	a := 0
	b := 0
	isA := 1
	for {
		if low == high {
			break
		}
		if piles[low] > piles[high] {
			if isA == 1 {
				a += piles[low]
			} else {
				b += piles[low]
			}
			low++
		} else {
			if isA == 1 {
				a += piles[high]
			} else {
				b += piles[high]
			}
			high--
		}
		isA = ^isA
	}
	if a > b {
		return true
	}
	return false
}
