package leetcode_go

/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start

// 回溯法
// func findTargetSumWays(nums []int, target int) int {
// 	ans := 0
// 	var solve func(sum int, i int)
// 	solve = func(sum, i int) {
// 		if i == len(nums)-1 {
// 			if sum+nums[i] == target {
// 				ans++
// 			}
// 			if sum-nums[i] == target {
// 				ans++
// 			}
// 			return
// 		}
// 		solve(sum+nums[i], i+1)
// 		solve(sum-nums[i], i+1)
// 	}
// 	solve(0, 0)
// 	return ans
// }

// 动态规划
func findTargetSumWays(nums []int, target int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	if (total-target)%2 != 0 || total-target < 0 {
		return 0
	}
	// 从数组中选择若干数，使其和为neg
	neg := (total - target) / 2
	//dp[i][j]从前i个里选择使和为j
	dp := make([][]int, len(nums)+1)
	for i, _ := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < neg+1; j++ {
			if j < nums[i] {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = dp[i][j] + dp[i][j-nums[i]]
			}
		}
	}
	return dp[len(nums)][neg]
}

// @lc code=end
