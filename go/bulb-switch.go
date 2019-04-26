package main

import (
	"fmt"
	"math"
)

func main() {
	test := []int{3, 4}
	for _, v := range test {
		fmt.Println(bulbSwitch(v))
	}
}

func bulbSwitch(n int) int {
	num := math.Sqrt(float64(n))
	return int(math.Floor(num))
}
