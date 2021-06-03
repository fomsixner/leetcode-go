package leetcode_go

/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 */

// 滑动窗口法
// 可以使用有序集合来判断
// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums) && j <= i+k; j++ {
			if nums[i]-nums[j] <= t && nums[i]-nums[j] >= -t {
				return true
			}
		}
	}
	return false
}

// @lc code=end
