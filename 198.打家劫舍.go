package leetcode_go

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	//当只有一间房时 dp[0] = nums[0]
	//当有2间房时 dp[0] = nums[0] dp[1] = max(nums[0], nums[1])
	//当房间数>2时 dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	// first 表示 dp[i-2], second表示dp[i-1]
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end
