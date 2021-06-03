package leetcode_go

/*
 * @lc app=leetcode.cn id=897 lang=golang
 *
 * [897] 递增顺序查找树
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

func increasingBST(root *TreeNode) *TreeNode {
	ans := &TreeNode{
		Left:  nil,
		Right: nil,
	}
	ptr := ans
	var inOrder func(*TreeNode)
	// 中序遍历
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		ptr.Right = root
		ptr = ptr.Right
		ptr.Left = nil
		inOrder(root.Right)
	}
	inOrder(root)
	return ans.Right
}

// @lc code=end
