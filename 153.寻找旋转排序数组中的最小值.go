package leetcode_go

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	// 二分查找法
	left, right := 0, len(nums)-1
	ans := nums[0]
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < ans {
			ans = nums[mid]
		}
		if nums[0] <= nums[mid] { //左边有序
			if nums[left] > nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右边有序
			right = mid - 1
		}
	}
	return ans
}

// @lc code=end
