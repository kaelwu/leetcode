package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("bb"))
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	newStr := insertWildcardtoStr(s)
	// 设定p数据为每一个字符回文子串的半径.
	p := make([]int, len(newStr))
	var mx, po, maxLen, position int
	for i := 1; i <= 2*len(s)+1; i++ {
		if mx > i {
			p[i] = min(mx-i+1, p[2*po-i])
		} else {
			p[i] = 1
		}
		for {
			if i-p[i] >= 0 && i+p[i] < len(newStr) && newStr[i-p[i]] == newStr[i+p[i]] {
				p[i]++
			} else {
				break
			}
		}
		if mx < i+p[i] {
			po = i
			mx = i + p[i]
		}
		tmp, res := max(maxLen, p[i])
		if !res {
			maxLen = tmp
			position = i
		}
	}
	maxLen--
	start := position - maxLen
	if start < 0 {
		start = 0
	}
	end := position - 2 + maxLen
	if end >= len(s) {
		end = len(s)
	}
	strOne := newStr[start:end]
	result := make([]byte, 0)
	for _, v := range strOne {
		if byte(v) == byte(35) {
			continue
		}
		result = append(result, byte(v))
	}
	return string(result)
}

func insertWildcardtoStr(s string) string {
	tmp := make([]byte, 0)
	tmp = append(tmp, byte(34))
	for i := 1; i <= 2*len(s); i += 2 {
		tmp = append(tmp, byte(35))
		tmp = append(tmp, s[i/2])
	}
	tmp = append(tmp, byte(35))
	tmp = append(tmp, byte(34))
	return string(tmp)
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func max(a, b int) (int, bool) {
	if a >= b {
		return a, true
	}
	return b, false
}
