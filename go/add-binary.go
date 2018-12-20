package main

import "fmt"

func main() {
	a := "1010"
	b := "1"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	enter := 0
	var now byte
	var long, short string
	if len(a) > len(b) {
		long = a
		short = b
	} else {
		long = b
		short = a
	}
	diffLen := len(long) - len(short)
	padding := make([]byte, 0)
	for i := 0; i < diffLen; i++ {
		padding = append(padding, byte(48))
	}
	short = string(padding) + short
	res := make([]byte, 0)
	for i := len(long) - 1; i >= 0; i-- {
		if int(long[i]) == 49 && int(short[i]) == 49 {
			now = byte(48 + enter)
			enter = 1
		} else if int(long[i]) == 48 && int(short[i]) == 48 {
			now = byte(48 + enter)
			enter = 0
		} else if enter == 1 && (long[i] == byte(49) || short[i] == byte(49)) {
			now = byte(48)
			enter = 1
		} else {
			now = byte(49)
			enter = 0
		}
		res = append([]byte{now}, res...)
	}
	if enter == 1 {
		res = append([]byte{49}, res...)
	}
	return string(res)
}
