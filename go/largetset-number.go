package main

import (
	"fmt"
	"sort"
)

type CmpStr struct {
	Str string
}

type CmpArr []CmpStr

func (c CmpArr) Len() int {
	return len(c)
}

func (c CmpArr) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c CmpArr) Less(i, j int) bool {
	tmp1 := fmt.Sprintf("%v%v", c[i].Str, c[j].Str)
	tmp2 := fmt.Sprintf("%v%v", c[j].Str, c[i].Str)
	return tmp1 > tmp2
}

func main() {
	fmt.Println(largestNumber([]int{81, 82, 8}))
}

func largestNumber(nums []int) string {
	newNums := make([]CmpStr, 0)
	for _, v := range nums {
		tmp := CmpStr{
			Str: fmt.Sprintf("%v", v),
		}
		newNums = append(newNums, tmp)
	}
	sort.Sort(CmpArr(newNums))
	res := ""
	for _, v := range newNums {
		res += v.Str
	}
	if string(res[0]) == "0" {
		return "0"
	}
	return res
}
