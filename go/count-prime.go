package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {
	count := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func isPrime(n int) bool {
	count := int(math.Sqrt(float64(n)))
	fmt.Println(n, count)
	for i := 2; i <= count; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
