package main

import "fmt"

func main() {
	fmt.Println(countAndSay(6))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)
	bs := []byte(str)
	old := string(bs[0])
	count := 1
	strNew := ""
	for _, v := range bs[1:] {
		tmp := string(v)
		if old == tmp {
			count++
		} else {
			strNew += fmt.Sprintf("%v%v", count, old)
			count = 1
			old = tmp
		}
	}
	strNew += fmt.Sprintf("%v%v", count, old)
	return strNew
}
