package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	tmp := head
	if head == nil {
		return head
	}
	newHead := &ListNode{
		Val:  head.Val,
		Next: nil,
	}
	res := newHead
	oldVal := head.Val
	for {
		if tmp == nil {
			break
		}
		val := tmp.Val
		if val != oldVal {
			newHead.Next = &ListNode{
				Val:  tmp.Val,
				Next: nil,
			}
			newHead = newHead.Next
		}
		oldVal = tmp.Val
		tmp = tmp.Next
	}
	return res
}
