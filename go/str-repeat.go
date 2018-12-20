package main

import "fmt"

func main() {
	len2 := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(len2)
}

type Node struct {
	Index int
	Len   int
}

func lengthOfLongestSubstring(s string) int {
	bs := []byte(s)
	m := make(map[string]int)
	var max, maxLen int
	for k, v := range bs {
		tmp := string(v)
		v1, ok := m[tmp]
		if ok {
			if max < v1 {
				max = v1
			}
		}
		if maxLen < k-max+1 {
			maxLen = k - max + 1
		}
		m[tmp] = k + 1
	}
	return maxLen
}
