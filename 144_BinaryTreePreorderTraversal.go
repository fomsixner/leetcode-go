package leetcode_go

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}



func preorderTraversal(root *TreeNode) []int {
	var res []int
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		res = append(res, node.Val);
		if node.Left != nil {
			preorder(node.Left)
		}
		if node.Right != nil {
			preorder(node.Right)
		}
	}
	if root != nil{
		preorder(root)
	}
	return res
}