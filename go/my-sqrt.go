package main

import "fmt"

func main() {
	test := []int{1, 4, 9, 8, 14, 19}
	for _, v := range test {
		fmt.Println(mySqrt(v))
	}
}

func mySqrt(x int) int {
	i := 1
	for {
		tmp := i * i
		if tmp < x {
			i++
		} else if tmp == x {
			return i
		} else {
			return i - 1
		}
	}
}
