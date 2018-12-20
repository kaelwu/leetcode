package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("a b       Hello wolrd"))
}

func lengthOfLastWord(s string) int {
	tmp := 0
	a := make([]int, 1)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if string(ch) == " " {
			if tmp != 0 {
				a = append(a, tmp)
			}
			tmp = 0
		} else {
			tmp++
		}
	}
	if tmp != 0 {
		a = append(a, tmp)
	}
	return a[len(a)-1]
}
