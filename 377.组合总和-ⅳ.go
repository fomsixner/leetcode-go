package leetcode_go

/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 */

// @lc code=start
// 动态规划
// dp[i] 表示总和为i的方案数
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1 // dp[0] = 1, 不选取任何元素
	for i := 1; i < target+1; i++ {
		for _, v := range nums {
			if v <= i {
				dp[i] += dp[i-v]
			}
		}
	}
	return dp[target]
}

// @lc code=end
