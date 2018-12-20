package main

import "fmt"

func main() {
	fmt.Println(reverseSlice([]int{1, 2, 3, 4, 5}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	a := make([]int, 0)
	tmp := copy(head)
	for {
		if tmp != nil {
			a = append(a, tmp.Val)
		} else {
			break
		}
		tmp = tmp.Next
	}
	b := reverseSlice(a)
	i := 0
	res := head
	for {
		if head != nil {
			head.Value = b[i]
		} else {
			break
		}
		head = head.Next
		i++
	}
	return res
}

func reverseSlice(s []int) []int {
	for i := 0; i < len(s)/2; i++ {
		tmp := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = tmp
	}
	return s
}
