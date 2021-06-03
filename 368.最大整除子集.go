package leetcode_go

import "sort"

/*
 * @lc app=leetcode.cn id=368 lang=golang
 *
 * [368] 最大整除子集
 */

// @lc code=start

// 递归法(超时)
// func maxSubset(nums []int, v int) []int {
// 	data := []int{}
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i]%v == 0 || v%nums[i] == 0 {
// 			data = append(data, nums[i])
// 		}
// 	}
// 	n := len(data)
// 	if n == 0 || n == 1 {
// 		return data
// 	}

// 	res := []int{}
// 	for i := 0; i < n; i++ {
// 		// 假设第i个数在子集中
// 		subset := []int{}
// 		for j := i + 1; j < n; j++ {
// 			temp := maxSubset(data[j:], data[i])
// 			if len(temp)+1 > len(subset) {
// 				subset = append(temp, data[i])
// 			}
// 		}
// 		if len(subset) > len(res) {
// 			res = subset
// 		}
// 	}
// 	return res
// }

// func largestDivisibleSubset(nums []int) []int {
// 	if len(nums) == 1 {
// 		return nums
// 	}
// 	ans := []int{}
// 	for i := 0; i < len(nums)-1; i++ {
// 		temp := maxSubset(nums[i+1:], nums[i])
// 		if len(temp)+1 > len(ans) {
// 			ans = append(temp, nums[i])
// 		}
// 	}
// 	return ans
// }

// 动态规划法
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([]int, n)
	// dp[i]表示以nums[i]为最大整数的最大子集
	for i := range dp {
		dp[i] = 1
	}

	maxSize, maxValue := 1, 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxSize, maxValue = dp[i], nums[i]
		}
	}
	if maxSize == 1 {
		return []int{nums[0]}
	}
	var res []int
	for i := n - 1; i >= 0 && maxSize > 0; i-- {
		if maxSize == dp[i] && maxValue%nums[i] == 0 {
			res = append(res, nums[i])
			maxSize--
			maxValue = nums[i]
		}
	}
	return res
}

// @lc code=end
