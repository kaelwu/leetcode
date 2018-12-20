package main

import "fmt"

func main() {
	fmt.Println(climbStairs(10))
}

func climbStairs(n int) int {
	a, b := 1, 1
	for {
		if n <= 0 {
			break
		}
		n--
		b += a
		a = b - a
	}
	return a
}
