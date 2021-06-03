package leetcode_go

import "math"

/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 */

// @lc code=start
func judgeSquareSum(c int) bool {
	a, b := 0, int(math.Sqrt(float64(c)))
	for a <= b {
		if a*a+b*b == c {
			return true
		}
		if a*a+b*b < c {
			a++
		} else {
			b--
		}
	}
	return false
}

// @lc code=end
