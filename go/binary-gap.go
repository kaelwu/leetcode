package main

import "fmt"

func main() {
	test := []int{4, 5, 6, 13, 22, 5049}
	for _, v := range test {
		fmt.Println(binaryGap(v))
	}
}

func binaryGap(N int) int {
	pos := 1
	var min, max, ret int
	for {
		if N == 0 {
			return ret
		}
		if N%2 == 0 {
			N = N / 2
		} else {
			if min == 0 && max == 0 {
				min = pos
				max = min
			} else {
				if pos-max >= ret {
					ret = pos - max
					min, max = max, pos
				} else {
					min = pos
					max = min
				}
			}
			N = (N - 1) / 2
		}
		pos++
	}
}
