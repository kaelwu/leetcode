package main

import "fmt"

func main() {
	test := [][]int{
		[]int{1, 2, 3, 4, 4, 3, 2, 1},
		[]int{1, 1, 1, 2, 2, 2, 3, 3},
		[]int{0, 0, 0, 1, 1, 1, 2, 2, 2},
	}
	for _, v := range test {
		fmt.Println(hasGroupsSizeX(v))
	}

}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func hasGroupsSizeX(deck []int) bool {
	if len(deck) <= 1 {
		return false
	}
	m := make(map[int]int)
	for _, v := range deck {
		if v1, ok := m[v]; ok {
			m[v] = v1 + 1
		} else {
			m[v] = 1
		}
	}
	nums := []int{}
	for _, v := range m {
		nums = append(nums, v)
	}
	if len(nums) == 1 {
		return true
	}
	g := 0
	for i := 1; i < len(nums); i++ {
		if g == 0 {
			g = gcd(nums[i], nums[0])
		} else {
			g = gcd(nums[i], g)
			if g == 1 {
				return false
			}
		}
	}
	if g == 0 {
		return false
	} else {
		return true
	}
}
