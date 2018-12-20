package main

import "fmt"

func main() {
	fmt.Println(step(5, 0))
}

func step(num, i int) int {
	if num == 1 {
		return i
	}
	if num%2 == 0 {
		num = min(num/2, num-1)
	} else {
		num = num - 1
	}

	return step(num, i+1)
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
