package leetcode_go

/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// func isBadVersion(version int) bool {
// 	return true
// }

func firstBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid
		}
		if left == right-1 {
			break
		}
	}
	if isBadVersion(left) {
		return left
	}
	return right
}

// @lc code=end
