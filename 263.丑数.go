package leetcode_go

/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		if n%2 == 0 {
			n /= 2
			continue
		} else if n%3 == 0 {
			n /= 3
			continue
		} else if n%5 == 0 {
			n /= 5
			continue
		}
		return false
	}
	return true
}

// @lc code=end

// 可以对 nn 反复除以 2,3,52,3,5，直到 nn 不再包含质因数 2,3,52,3,5。
// 若剩下的数等于 11，则说明 nn 不包含其他质因数，是丑数；否则，说明 nn 包含其他质因数，不是丑数。
// var factors = []int{2, 3, 5}

// func isUgly(n int) bool {
//     if n <= 0 {
//         return false
//     }
//     for _, f := range factors {
//         for n%f == 0 {
//             n /= f
//         }
//     }
//     return n == 1
// }
