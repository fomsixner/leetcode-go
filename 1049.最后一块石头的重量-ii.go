package leetcode_go

/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
// 0-1背包问题变种
// 从stones数组中选择，凑成不大于sum/2的最大数
func lastStoneWeightII(stones []int) int {
	sum, n := 0, len(stones)
	for _, v := range stones {
		sum += v
	}
	half := sum / 2
	dp := make([]int, half+1)
	for i := 0; i < n; i++ {
		for j := half; j >= stones[i]; j-- {
			if dp[j-stones[i]]+stones[i] > dp[j] {
				dp[j] = dp[j-stones[i]] + stones[i]
			}
		}
	}
	return sum - 2*dp[half]
}

// @lc code=end
