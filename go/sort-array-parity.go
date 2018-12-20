package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func sortArrayByParityII(A []int) []int {
	odd := make([]int, 0)
	even := make([]int, 0)
	for _, v := range A {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	for i := 0; i < len(A); i++ {
		if i%2 == 0 {
			A[i] = even[i/2]
		} else {
			A[i] = odd[(i-1)/2]
		}
	}
	return A
}

func sortArrayByParityII(A []int) []int {
	oldIndex := -1
	odd := make([]int, 0)
	even := make([]int, 0)
	for k, v := range A {
		if k%2 != v%2 {
			if v%2 == 1 {
				odd = append(odd, k)
			} else {
				even = append(even, k)
			}
		}
	}
	for i := 0; i < len(odd); i++ {
		A[odd[i]], A[even[i]] = A[even[i]], A[odd[i]]
	}
	return A
}

func sortArrayByParityII(A []int) []int {
	j := 1
	for i := 0; i < len(A); i += 2 {
		if A[i]%2 != 0 {
			for {
				if A[j]%2 == 0 {
					A[i], A[j] = A[j], A[i]
					break
				}
				j += 2
			}
		}
	}
	return A
}
