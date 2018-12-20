package main

import "fmt"

func main() {
	fmt.Println(trailingZeroes(5))
}

func trailingZeroes(n int) int {
	sum := 0
	for {
		if n <= 0 {
			break
		}
		sum += n / 5
		n /= 5
	}
	return sum
}
