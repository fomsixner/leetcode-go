package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumNumbers(root *TreeNode) int {
	var dfs func(node *TreeNode, sum int) int
	dfs = func(node *TreeNode, sum int) int {
		if node == nil {
			return 0
		}
		sum = sum * 10 + node.Val
		if node.Left == nil && node.Right == nil {
			return sum
		} else {
			return dfs(node.Left, sum) + dfs(node.Right, sum)
		}
	}
	return dfs(root, 0)
}