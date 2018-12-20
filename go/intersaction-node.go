package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	for n1 := headA; n1 != nil; n1 = n1.Next {
		for n2 := headB; n2 != nil; n2 = n2.Next {
			if n1 == n2 {
				return n1
			}
		}
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	n1 := headA
	n2 := headB
	for n1 != n2 {
		if n1 == nil {
			n1 = headB
		} else {
			n1 = n1.Next
		}
		if n2 == nil {
			n2 = headA
		} else {
			n2 = n2.Next
		}
	}
	return n1
}
