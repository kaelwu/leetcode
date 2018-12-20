package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4}))
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
		tmp int
	)

	for i := 0; i < l-1; i++ {
		p[i] = v[i+1] - v[i]
	}

	for _, t := range p {
		if tmp+t > 0 {
			tmp += t
		} else {
			tmp = 0
		}
		if tmp > max {
			max = tmp
		}
	}

	return max
}

func maxProfit(v []int) int {
	i := 0
	j := len(v) - 1
	for {
		if v[i]-v[j] < 0 {
			i++
		}
	}
}
