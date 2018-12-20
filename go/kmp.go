package main

import "fmt"

func main() {
	fmt.Println(getKMP("ababc", "abcd"))
}

func getNext(p []int) []int {
	result := make([]int, len(p))
	result[0] = -1
	j, k := 0, -1
	for j < len(p)-1 {
		if k == -1 || p[j] == p[k] {
			j = j + 1
			k = k + 1
			result[j] = k
		} else {
			k = result[k]
		}
	}
	return result
}

func getKMP(a, b string) int {
	var i, j int
	next := make([]int, len(b))
	for i < len(a) && j < len(b) {
		if j == -1 || a[i] == b[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(b) {
		return i - j
	}
	return -1
}
