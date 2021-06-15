package leetcode_go

/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 */

// @lc code=start
// 二分查找法
func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	var mid int
	for left < right {
		mid = (left + right) / 2
		if arr[mid] < arr[mid-1] {
			right = mid
		} else if arr[mid] < arr[mid+1] {
			left = mid
		} else {
			break
		}
	}
	return mid
}

// @lc code=end
