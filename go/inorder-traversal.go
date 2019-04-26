package main

import "fmt"

func main() {
	root := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(inorderTraversal(&root))
}

var res []int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res = []int{}
	printNode(root)
	return res
}

func printNode(node *TreeNode) {
	if node.Left != nil {
		printNode(node.Left)
	}
	res = append(res, node.Val)
	if node.Right != nil {
		printNode(node.Right)
	}
}
