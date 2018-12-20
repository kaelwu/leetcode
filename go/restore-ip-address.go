package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(restoreIPAddresses("25525511135"))
}

func restoreIPAddresses(s string) []string {
	res := make([]string, 0)
	findIP(s, 0, 0, "", res)
	return res
}

func findIP(s string, f, idx int, ip string, res []string) {
	if idx == 3 {
		if len(s)-1-f < 3 {
			if string(s[f]) == "0" && f != len(s)-1 {
				return
			}
			num, _ := strconv.Atoi(s[f:])
			if num < 256 {
				ip += s[f:]
				res = append(res, ip)
			}
		}
	} else {
		for i := 1; i <= 3; i++ {
			if f+i > len(s) {
				break
			}
			num, _ := strconv.Atoi(s[f : f+i])
			if num < 256 {
				findIP(s, f+i, idx+1, ip+s[f:f+i]+".", res)
			}
			if string(s[f]) == "0" && i == 1 {
				break
			}
		}
	}
}
