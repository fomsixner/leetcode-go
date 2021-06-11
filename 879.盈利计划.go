package leetcode_go

/*
 * @lc app=leetcode.cn id=879 lang=golang
 *
 * [879] 盈利计划
 */

// @lc code=start
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	const mod = 1e9 + 7
	gl := len(profit)
	dp := make([][][]int, gl+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, minProfit+1)
		}
	}
	dp[0][0][0] = 1
	for i := 0; i < gl; i++ {
		for j := 0; j < n+1; j++ {
			for k := 0; k < minProfit+1; k++ {
				if group[i] <= j {
					dp[i+1][j][k] = (dp[i][j][k] + dp[i][j-group[i]][max(0, k-profit[i])]) % mod
				} else {
					dp[i+1][j][k] = dp[i][j][k]
				}
			}
		}
	}
	res := 0
	for _, d := range dp[gl] {
		res = (res + d[minProfit]) % mod
	}
	return res
}

// @lc code=end
