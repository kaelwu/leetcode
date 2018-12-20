package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}

func plusOne(digits []int) []int {
	length := len(digits)
	nd := make([]int, 1)
	nd = append(nd, digits...)
	i := length
	for {
		if nd[i]+1 < 10 {
			nd[i] = nd[i] + 1
			break
		} else {
			nd[i] = 0
			i--
		}
	}
	if nd[0] == 0 {
		return nd[1:]
	}
	return nd
}
