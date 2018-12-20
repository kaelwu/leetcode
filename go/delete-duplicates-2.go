package main

import "fmt"

func main() {
	a1 := &ListNode{
		Val: 1,
	}
	a2 := &ListNode{
		Val: 2,
	}
	a3 := &ListNode{
		Val: 2,
	}
	a4 := &ListNode{
		Val: 3,
	}
	a5 := &ListNode{
		Val: 3,
	}
	a6 := &ListNode{
		Val: 4,
	}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	a5.Next = a6
	fmt.Println(deleteDuplicates(a1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	new := head
	m := make(map[int]int, 0)
	for {
		if new == nil {
			break
		}
		if v, ok := m[new.Val]; ok {
			m[new.Val] = v + 1
		} else {
			m[new.Val] = 1
		}
		new = new.Next
	}
	fmt.Println(m)
	current := &ListNode{}
	res := current
	for {
		if head == nil {
			break
		}
		if m[head.Val] == 1 {
			tmp := &ListNode{
				Val:  head.Val,
				Next: nil,
			}
			current.Next = tmp
			current = current.Next
		}
		head = head.Next
	}
	return res.Next
}
