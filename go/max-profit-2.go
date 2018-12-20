package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 6, 3, 2, 1, 0}))
}

func maxProfit(v []int) int {
	var (
		l = len(v)
	)

	if l < 2 {
		return 0
	}

	var (
		p   = make([]int, l-1)
		max int
	)

	for i := 0; i < l-1; i++ {
		p[i] = v[i+1] - v[i]
	}

	for _, t := range p {
		if t > 0 {
			max += t
		}
	}

	return max
}
