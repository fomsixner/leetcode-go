package leetcode_go

/*
 * @lc app=leetcode.cn id=523 lang=golang
 *
 * [523] 连续的子数组和
 */

// @lc code=start
func checkSubarraySum(nums []int, k int) bool {
	pre := make([]int, len(nums)+1)
	count := 0
	for i, _ := range nums {
		pre[i+1] = pre[i] + nums[i]
		if nums[i] == 0 {
			count++
		} else {
			count = 0
		}
		if count >= 2 {
			return true
		}
	}
	for i := 2; i < len(pre); i++ {
		for left := 0; left < i-1; left++ {
			if (pre[i]-pre[left])%k == 0 {
				return true
			}
			if (pre[i] - pre[left]) < k {
				break
			}
		}
	}
	return false
}

// @lc code=end
