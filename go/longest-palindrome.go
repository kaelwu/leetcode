package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("bb"))
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	new_str := insertWildcardtoStr(s)
	// 设定p数据为每一个字符回文子串的半径.
	p := make([]int, len(new_str))
	var mx, po, max_len, position int
	for i := 1; i <= 2*len(s)+1; i++ {
		if mx > i {
			p[i] = min(mx-i+1, p[2*po-i])
		} else {
			p[i] = 1
		}
		for {
			if i-p[i] >= 0 && i+p[i] < len(new_str) && new_str[i-p[i]] == new_str[i+p[i]] {
				p[i]++
			} else {
				break
			}
		}
		if mx < i+p[i] {
			po = i
			mx = i + p[i]
		}
		tmp, res := max(max_len, p[i])
		if !res {
			max_len = tmp
			position = i
		}
	}
	max_len--
	fmt.Println(max_len, position)
	start := position - max_len
	if start < 0 {
		start = 0
	}
	end := position - 2 + max_len
	if end >= len(s) {
		end = len(s)
	}
	fmt.Println(start, end)
	str_one := new_str[start:end]
	result := make([]byte, 0)
	for _, v := range str_one {
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
	} else {
		return a
	}
}

func max(a, b int) (int, bool) {
	if a >= b {
		return a, true
	} else {
		return b, false
	}
}
