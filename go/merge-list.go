package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n1 := &ListNode{}
	n2 := &ListNode{}
	new := mergeTwoLists(n1, n2)
	for {
		if new == nil {
			return
		}
		fmt.Println(new.Val)
		new = new.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var head *ListNode

	if l1 == nil && l2 == nil {
		return l1
	}
	if l1 != nil && l2 == nil {
		head = l1
		l1 = l1.Next
	}
	if l2 != nil && l1 == nil {
		head = l2
		l2 = l2.Next
	}
	if l2 != nil && l1 != nil {
		if l1.Val < l2.Val {
			head = l1
			l1 = l1.Next
		} else {
			head = l2
			l2 = l2.Next
		}
	}

	var pre *ListNode = head
	var node *ListNode
	for l1 != nil && l2 != nil {

		if l1.Val < l2.Val {
			node = l1
			l1 = l1.Next
		} else {
			node = l2
			l2 = l2.Next
		}
		pre.Next = node
		pre = node

	}
	if l1 != nil {
		node = l1
		pre.Next = node
		pre = node
		l1 = l1.Next

	}
	if l2 != nil {
		node = l2
		pre.Next = node
		pre = node
		l2 = l2.Next

	}
	return head

}
