package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numUniqueEmails([]string{"1233.123+123@asd.com"}))
}

func numUniqueEmails(emails []string) int {
	m := make(map[string]bool)
	count := 0
	for _, v := range emails {
		newstrs := strings.Split(v, "@")
		news := strings.Split(newstrs[0], "+")
		p := news[0]
		prefix := strings.Replace(p, ".", "", 0)
		newstr := prefix + "@" + newstrs[1]
		if _, ok := m[newstr]; ok {
			continue
		} else {
			m[newstr] = true
			count++
		}
	}
	fmt.Println(m)
	return count
}

func replace(a rune) rune {
	b := []byte(".")
	if a == rune(b[0]) {
		return -1
	}
	return a
}
