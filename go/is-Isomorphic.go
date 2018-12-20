package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("foo", "baa"))
}

func isIsomorphic(s string, t string) bool {
	length := len(s)
	n1 := make([]byte, length)
	n2 := make([]byte, length)
	// 设置映射关系
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)
	for i := 0; i < length; i++ {
		m1[s[i]] = t[i]
		m2[t[i]] = s[i]
	}
	for i := 0; i < length; i++ {
		n1[i] = m1[s[i]]
		n2[i] = m2[t[i]]
	}
	if string(n1) == t && string(n2) == s {
		return true
	}
	return false
}
