package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

var ans int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans = 0
	deep(root)
	return ans
}

func deep(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}
	leftDeep := deep(root.Left)
	rightDeep := deep(root.Right)
	ans = max(ans, leftDeep+rightDeep)
	return 1 + max(leftDeep, rightDeep)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
