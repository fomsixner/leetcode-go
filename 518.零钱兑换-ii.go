package leetcode_go

/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
// func change(amount int, coins []int) int {
// 	n := len(coins)
// 	dp := make([][]int, n+1)
// 	for i := range dp {
// 		dp[i] = make([]int, amount+1)
// 	}
// 	dp[0][0] = 1
// 	for i, coin := range coins {
// 		for j := 0; j < amount+1; j++ {
// 			dp[i+1][j] = dp[i][j]
// 			for k := 1; k*coin <= j; k++ {
// 				dp[i+1][j] += dp[i][j-k*coin]
// 			}
// 		}
// 	}
// 	return dp[n][amount]
// }

// 内存优化版本
// func change(amount int, coins []int) int {
// 	dp := make([]int, amount+1)
// 	dp[0] = 1
// 	for _, coin := range coins {
// 		for j := amount; j >= 0; j-- {
// 			for k := 1; k*coin <= j; k++ {
// 				dp[j] += dp[j-k*coin]
// 			}
// 		}
// 	}
// 	return dp[amount]
// }

// 继续优化版本, 上一个版本中重复计算很多
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j < amount+1; j++ {
			dp[j] += dp[j-coin]
		}
	}
	return dp[amount]
}

// @lc code=end
