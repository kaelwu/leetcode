package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uncommonFromSentences("s z z z s", "s z etj"))
}

// 把字符串拼接起来出现一次的统计.
func uncommonFromSentences(A string, B string) []string {
	a := strings.Split(A, " ")
	b := strings.Split(B, " ")
	am := make(map[string]int)
	bm := make(map[string]int)
	for _, v := range a {
		if v1, ok := am[v]; !ok {
			am[v] = 1
		} else {
			am[v] = v1 + 1
		}
	}
	for _, v := range b {
		if v2, ok := bm[v]; !ok {
			bm[v] = 1
		} else {
			bm[v] = v2 + 1
		}
	}
	res := []string{}
	for _, v := range a {
		if am[v] > 1 {
			continue
		}
		if _, ok := bm[v]; !ok {
			res = append(res, v)
		}
	}
	for _, v := range b {
		if bm[v] > 1 {
			continue
		}
		if _, ok := am[v]; !ok {
			res = append(res, v)
		}
	}

	return res
}
