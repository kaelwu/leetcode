package main

import "fmt"

func main() {
	fmt.Println(strStr("mississippi", "issipp"))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	var i, j, k int
	bs1 := []byte(haystack)
	bs2 := []byte(needle)
	for {
		if i >= len(bs1) || j == len(bs2) {
			break
		}
		tmp1 := string(bs1[i])
		tmp2 := string(bs2[j])
		if tmp1 != tmp2 {
			j = 0
			k++
			i = k
		} else {
			j++
			i++
		}

	}
	if j == len(bs2) {
		return i - j
	}
	return -1
}
