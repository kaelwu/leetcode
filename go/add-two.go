package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a1 := &ListNode{
		Val: 2,
	}
	a2 := &ListNode{
		Val: 4,
	}
	a3 := &ListNode{
		Val: 3,
	}
	a1.Next = a2
	a2.Next = a3
	b1 := &ListNode{
		Val: 5,
	}
	b2 := &ListNode{
		Val: 6,
	}
	b3 := &ListNode{
		Val: 7,
	}
	b1.Next = b2
	b2.Next = b3
	addTwoNumbers(a3, b3)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tmp1, tmp2, carry int
	times1 := 1
	times2 := 1
	nums := make([]*ListNode, 0)
	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		if l1 != nil {
			tmp1 = l1.Val
			l1 = l1.Next
		} else {
			tmp1 = 0
		}
		if l2 != nil {
			tmp2 = l2.Val
			l2 = l2.Next
		} else {
			tmp2 = 0
		}
		num := tmp1 + tmp2 + carry
		if num >= 10 {
			num = num - 10
			carry = 1
		} else {
			carry = 0
		}
		fmt.Println(num)
		nums = append(nums, &ListNode{Val: num})
		times1 = times1 * 10
		times2 = times2 * 10
	}
	for k, v := range nums {
		if k+1 <= len(nums)-1 {
			v.Next = nums[k+1]
		} else {
			v.Next = nil
		}
	}
	return nums[0]
}
