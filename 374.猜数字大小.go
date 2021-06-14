package leetcode_go

import "math"

/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] 猜数字大小
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    left, right := 1, math.MaxInt32
	mid := (left + right)/2
	for left < right {
		if guess(mid) == 0 {
			break 
		} else if guess(mid) < 0 {
			right = mid -1
		} else {
			left = mid + 1
		}
		mid = (left + right)/2
	}
	return mid 
}
// @lc code=end

