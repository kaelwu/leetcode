package main

import "fmt"

func main() {
	fmt.Println(isValid("{[}]"))
}

func isValid(s string) bool {
	bs := []byte(s)
	res := true
	chars := make([]string, 0)
	for _, v := range bs {
		tmp := string(v)
		if tmp == ")" || tmp == "]" || tmp == "}" {
			if len(chars) >= 1 {
				check := chars[len(chars)-1]
				if (tmp == ")" && check == "(") || (tmp == "]" && check == "[") || (tmp == "}" && check == "{") {
					chars = chars[:len(chars)-1]
				} else {
					chars = append(chars, tmp)
				}
			} else {
				chars = append(chars, tmp)
			}
		} else {
			chars = append(chars, tmp)
		}
	}
	if len(chars) > 0 {
		return false
	}
	return res
}
