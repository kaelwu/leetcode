package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func majorityElement(nums []int) int {
	tmp := 0
	count := 0
	for _, v := range nums {
		if count == 0 {
			tmp = v
		}
		if v == tmp {
			count++
		} else {
			count--
		}
	}
	return tmp
}
