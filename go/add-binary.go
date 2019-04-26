package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "1010"
	b := "1"
	fmt.Println(addBinary(a, b))
	fmt.Println(addBinaryTwo(a, b))
}

func addBinaryTwo(a, b string) string {
	var res string
	m := len(a) - 1
	n := len(b) - 1
	carry := 0
	var tmpOne, tmpTwo int
	for {
		if m < 0 && n < 0 {
			if carry == 1 {
				return "1" + res
			}
			return res
		}
		if m >= 0 {
			tmpOne = int(a[m])
			m = m - 1
		} else {
			tmpOne = 0
		}
		if n >= 0 {
			tmpTwo = int(b[n])
			n = n - 1
		} else {
			tmpTwo = 0
		}
		sum := tmpOne + tmpTwo + carry
		res = strconv.Itoa(sum%2) + res
		carry = sum / 2
	}
	return res
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
