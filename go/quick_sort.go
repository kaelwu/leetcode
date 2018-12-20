package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{2, 4, 1, 1, 6, 9, 8, 10, 2, 4, 7, 23, 11}
	quick_sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	return
}

func findPivot(arr []int) int {
	count := len(arr)
	if count <= 0 {
		return 0
	}
	return rand.Intn(count - 1)
}

func random_partition(aris []int, begin int, end int) int {
	// 随机化哨兵.
	random := rand.Intn(end-begin) + begin
	fmt.Println(random, end)
	aris[end], aris[random] = aris[random], aris[end]
	pvalue := aris[end]
	i := begin - 1
	for j := begin; j < end; j++ {
		if aris[j] <= pvalue {
			i++
			aris[j], aris[i] = aris[i], aris[j]
		}
	}
	aris[i+1], aris[end] = aris[end], aris[i+1]
	return i + 1
}

func quick_sort(aris []int, begin int, end int) {
	if begin < end {
		mid := random_partition(aris, begin, end)
		quick_sort(aris, begin, mid-1)
		quick_sort(aris, mid+1, end)
	}
}
