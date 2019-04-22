package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(bagOfTokensScore([]int{26}, 51))
}

func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	// 用能量换小的，用分数换大的
	i := 0
	j := len(tokens) - 1
	count := 0
	max_count := 0
	for {
		if i > j {
			break
		}
		if P < tokens[i] && count == 0 {
			return count
		}
		if P >= tokens[i] {
			count++
			P = P - tokens[i]
			i++
			max_count = max(max_count, count)
			continue
		} else {
			P += tokens[j]
			j--
			count--
		}
	}
	return max_count
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
