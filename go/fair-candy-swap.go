package main

import "fmt"

type Test struct {
	A []int
	B []int
}

func main() {
	test := []Test{
		Test{
			A: []int{1, 1},
			B: []int{2, 2},
		},
		Test{
			A: []int{1, 2, 5},
			B: []int{2, 4},
		},
	}
	for _, v := range test {
		fmt.Println(fairCandySwap(v.A, v.B))
	}
}

func fairCandySwap(A []int, B []int) []int {
	var totalA, totalB int
	for _, v := range A {
		totalA += v
	}
	for _, v := range B {
		totalB += v
	}
	need := (totalA + totalB) / 2
	diff := need - totalA
	for _, v1 := range A {
		for _, v2 := range B {
			if v2-v1 == diff {
				return []int{v1, v2}
			}
		}
	}
	return []int{}
}
