package leetcode_go

/*
 * @lc app=leetcode.cn id=1744 lang=golang
 *
 * [1744] 你能在你最喜欢的那天吃到你最喜欢的糖果吗？
 */

// @lc code=start
func canEat(candiesCount []int, queries [][]int) []bool {
	// 糖果的前缀和
	pre := make([]int, len(candiesCount)+1)
	for i := 0; i < len(candiesCount); i++ {
		pre[i+1] = pre[i] + candiesCount[i]
	}
	ans := make([]bool, len(queries))
	for i, q := range queries {
		// 每天最少吃一颗糖果，最多吃q[2]颗糖果
		// 累计吃的糖果数目必须满足以下条件
		ans[i] = pre[q[0]+1] > 1*q[1] && pre[q[0]] < (q[2]*q[1]+q[2])
	}
	return ans
}

// @lc code=end
