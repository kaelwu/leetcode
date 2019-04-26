package main

import "fmt"

func main() {
	tests := []int{38, 122}
	for _, v := range tests {
		fmt.Println(addDigit(v))
	}
}

func addDigit(n int) int {
	if n/10 >= 1 {
		num := 0
		for {
			num += n % 10
			n = int(n / 10)
			if n == 0 {
				return addDigit(num)
			}
		}
	}
	return n
}
