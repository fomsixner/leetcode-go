package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	res := 0
	queue = append(queue, root)
	for len(queue) != 0 {
		cur := queue[0]
		res++
		queue = queue[1:]
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return res
}

//方法二: 根据完全二叉树的特性，从左子树找树的高度
//先计算树的高度height，然后计算右子树的高度
//1.如果右子树的高度等于height-1，说明左子树是完美二叉树（如下图所示），可以通过公式(2^(height-1))-1计算即可，
//  不需要全部遍历，然后再通过递归的方式计算右子树……，
//2.如果右子树的高度不等于height-1，说明右子树是完美二叉树（如下图所示），
//  只不过比上面那种少了一层，也就是height-2，也可以通过公式(2^(height-2))-1计算，然后再通过递归的方式计算左子树……，
