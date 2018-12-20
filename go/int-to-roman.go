package main

import "fmt"

var ROMAN_MAP = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func main() {
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	nums := make([]int, 0)
	i := 1
	for {
		if num >= 1 {
			nums = append(nums, num%10*i)
			num = int(num / 10)
			i = i * 10
		} else {
			break
		}
	}
	new_nums := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		new_nums = append(new_nums, nums[i])
	}

	return "123"
}
