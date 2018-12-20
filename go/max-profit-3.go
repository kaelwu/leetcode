package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 4, 2}))
	fmt.Println(maxProfit([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
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
		max = make([]int, 2)
	)

	for i := 0; i < l-1; i++ {
		p[i] = v[i+1] - v[i]
	}

	tmp := 0
	for _, t := range p {
		if t > 0 {
			tmp += t
		} else {
			fmt.Println(max)
			if tmp > max[0] {
				max[1] = max[0]
				max[0] = tmp
			} else if tmp > max[1] {
				max[1] = tmp
			}
			tmp = 0
		}
	}
	fmt.Println(max, tmp)
	// 排序到最后还需要检查一次
	if tmp > 0 {
		if tmp > max[0] {
			max[1] = max[0]
			max[0] = tmp
		} else if tmp > max[1] {
			max[1] = tmp
		}
	}

	return max[0] + max[1]
}
