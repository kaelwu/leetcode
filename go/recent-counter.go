package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type RecentCounter struct {
	t []int
}

func Constructor() RecentCounter {
	r := RecentCounter{}
	r.t = make([]int, 0)
}

func (this *RecentCounter) Ping(t int) int {
	this.t = append(this.t, t)
	count := 0
	for _, v := range this.t {
		if v >= t-3000 && v <= t {
			count++
		}
	}
	return count
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
