package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Result struct {
	List []int
}

func main() {
	a := &TreeNode{2, &TreeNode{5, nil, nil}, &TreeNode{2, nil, nil}}
	fmt.Println(inorderTraversal(a))
}

func inorderTraversal(root *TreeNode) []int {
	res := Result{}
	res.traversal(root)
	return res.List
}

func (res *Result) traversal(node *TreeNode) {
	if node != nil {
		res.traversal(node.Left)
		res.List = append(res.List, node.Val)
		res.traversal(node.Right)
	}
}
