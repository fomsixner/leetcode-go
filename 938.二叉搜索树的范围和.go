package leetcode_go

/*
 * @lc app=leetcode.cn id=938 lang=golang
 *
 * [938] 二叉搜索树的范围和
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
func rangeSumBST(root *TreeNode, low int, high int) int {
	// 中序遍历得到递增序列
	ans, flag := 0, 0
	var inOrder func(*TreeNode)
	inOrder = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		inOrder(tn.Left)
		if tn.Val == low {
			flag = 1
		}
		if flag == 1 {
			ans += tn.Val
		}
		if tn.Val == high {
			flag = 0
		}
		inOrder(tn.Right)
	}
	inOrder(root)
	return ans
}

// @lc code=end
