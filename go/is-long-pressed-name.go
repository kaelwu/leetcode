package main

import (
	"fmt"
)

func main() {
	fmt.Println(isLongPressedName("alex", "aaleex"))
}

func isLongPressedName(name string, typed string) bool {
	if name == typed {
		return true
	}
	if len(name) > len(typed) || name[0] != typed[0] || name[len(name)-1] != typed[len(typed)-1] {
		return false
	}

	i := 0
	j := 0
	prev := 0
	for {
		if i >= len(name) || j >= len(typed) {
			break
		}
		if name[i] == typed[j] {
			prev = i
			i++
			j++
		} else if typed[j] == name[prev] {
			j++
		} else {
			return false
		}
	}
	return true
}
