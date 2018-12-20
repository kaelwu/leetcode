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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	i := 1
	return count(root, i)
}

func count(t *TreeNode, i int) int {
	if t.Left != nil && t.Right != nil {
		return min(count(t.Left, i+1), count(t.Right, i+1))
	}
	if t.Left != nil {
		return count(t.Left) + 1
	}
	if t.Right != nil {
		return count(t.Right) + 1
	}
	return i
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
