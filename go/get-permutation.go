package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getPermutation(4, 10))
}

func getPermutation(n int, k int) string {
	f := make([]int, n)
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	f[0] = 1
	for i := 1; i < n; i++ {
		f[i] = f[i-1] * i
	}
	k--
	res := ""
	for i := n; i >= 1; i-- {
		j := k / f[i-1]
		k %= f[i-1]
		res += strconv.Itoa(intArr[j])
		intArr = append(intArr[:j], intArr[j+1:]...)
	}
	return res
}
