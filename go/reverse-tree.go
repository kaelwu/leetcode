package main

import (
	"fmt"
	"math"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	test := []string{
		"1 2 3 4 5 6",
		"4 2 7 # 3 6 9",
	}
	for _, v := range test {
		fmt.Println(buildTree(v))
	}
}

func buildTree(s string) string {
	arr := strings.Split(s, " ")
	length := len(arr)
	if length == 0 || length == 1 {
		return s
	}
	layer := make(map[int][]string)
	for i := 0; i < len(arr); i++ {
		l := int(math.Floor(math.Log2(float64(i + 1))))
		if v, ok := layer[l]; ok {
			v = append(v, arr[i])
			layer[l] = v
		} else {
			v = []string{arr[i]}
			layer[l] = v
		}
	}
	allLayerNum := int(math.Floor(math.Log2(float64(length)))) + 1
	var res string
	for i := 0; i < allLayerNum; i++ {
		v := layer[i]
		for j := len(v) - 1; j >= 0; j-- {
			tmp := v[j] + " "
			res += tmp
		}
	}
	return res
}
