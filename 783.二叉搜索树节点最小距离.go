package leetcode_go

/*
 * @lc app=leetcode.cn id=783 lang=golang
 *
 * [783] 二叉搜索树节点最小距离
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 二叉搜索树中序遍历得到的序列为有序的
func minDiffInBST(root *TreeNode) int {
	values := []int{}
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			inorder(root.Left)
		}
		values = append(values, root.Val)
		if root.Right != nil {
			inorder(root.Right)
		}
	}
	inorder(root)
	ans := values[1] - values[0]
	for i := 2; i < len(values); i++ {
		if values[i]-values[i-1] < ans {
			ans = values[i] - values[i-1]
		}
	}
	return ans
}

// @lc code=end
