package main

import "fmt"

var ROMAN_MAP = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func main() {
	fmt.Println(romanToInt("XXVII"))
}

func romanToInt(s string) int {
	nums := make([]int, 0)
	for _, v := range s {
		num := ROMAN_MAP[string(v)]
		nums = append(nums, num)
	}
	res := 0
	for k := range nums {
		if k+1 < len(nums) && nums[k] < nums[k+1] {
			res -= nums[k]
		} else {
			res += nums[k]
		}
	}
	return res
}
