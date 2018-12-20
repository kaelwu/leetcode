package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	i := 0
	if root == nil {
		return 0
	}
	count(root, i)
}

func count(t *TreeNode, i int) int {
	if t.Left != nil && t.Right == nil {
		count(t.Left, i+1)
	} else if t.Left == nil && t.Right != nil {
		count(t.Left, i+1)
	} else if t.Left == nil && t.Left == nil {
		return i
	} else {
		return max(count(t.Left, i+1), count(t.Right, i+1))
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
