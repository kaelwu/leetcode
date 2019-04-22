package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	i := 0
	m := make(map[int]*ListNode)
	for {
		if head == nil {
			break
		}
		m[i] = head
		head = head.Next
		i++
	}
	mid := int(i/2) + 1
	if i == 0 {
		return nil
	}
	return m[mid]
}

func middleNode(head *ListNode) *ListNode {
	low := head
	fast := head
	for {
		if fast != nil && fast.Next != nil {
			low = low.Next
			fast = fast.Next.Next
		} else {
			break
		}
	}
	return low
}
