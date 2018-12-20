package main

import (
	"fmt"
)

func main() {
	a1 := &TreeNode{
		Val: 0,
	}
	b1 := &TreeNode{
		Val: 0,
	}
	fmt.Println(isSameTree(a1, b1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		if p == q {
			return true
		}
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
