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

var res map[int][]int

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	levelPrint(root, 1)
	fmt.Println(res)
	res2 := make([][]int, 0)
	for i := 1; i <= len(res); i++ {
		res2 = append(res2, res[i])
	}
	return res2
}

func levelPrint(root *TreeNode, level int) {
	if _, ok := res[level-1]; ok {
		res[level-1] = append(res[level-1], root.Val)
	} else {
		res[level-1] = []int{root.Val}
	}
	if root.Left != nil {
		levelPrint(root.Left, level+1)
	}
	if root.Right != nil {
		levelPrint(root.Right, level+1)
	}
}
