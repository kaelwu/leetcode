package main

import "fmt"

type S struct {
	N []int
	T int
}

func main() {
	test := []S{
		S{
			N: []int{5, 7, 7, 8, 8, 10},
			T: 8,
		},
		S{
			N: []int{5, 7, 7, 8, 8, 10},
			T: 6,
		},
		S{
			N: []int{1, 2, 3},
			T: 3,
		},
	}
	for _, v := range test {
		fmt.Println(searchRange(v.N, v.T))
	}
}

func searchRange(nums []int, target int) []int {
	l := len(nums) - 1
	var mid int
	min := -1
	max := -1
	i := 0
	for {
		if i > l {
			break
		}
		mid = i + (l-i)/2
		if nums[mid] > target {
			l = l - 1
			continue
		} else if nums[mid] > target {
			i = i + 1
			continue
		} else {
			i = mid
			l = mid
			fmt.Println(i, l)
			for {
				if i > 0 && nums[i-1] == nums[i] {
					min = i - 1
					i--
				} else {
					break
				}
			}
			for {
				if l < len(nums)-1 && nums[l+1] == nums[l] {
					max = l + 1
					l++
				} else {
					break
				}
			}
			return []int{i, l}
		}

	}
	return []int{min, max}
}
