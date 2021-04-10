package leetcode_go

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start

func generateParenthesis(n int) []string {
	var res []string
	var backtrack func(str string, left int, right int)
	backtrack = func(str string, left int, right int) {
		// left, right 当前左右括号数
		if left == n && right == n {
			res = append(res, str)
			return
		}
		if left == right {
			backtrack(str+"(", left+1, right)
		} else if left > right {
			if left < n {
				backtrack(str+"(", left+1, right)
			}
			backtrack(str+")", left, right+1)
		}
		// left < right 为失败的情况
	}
	backtrack("", 0, 0)
	return res
}

// @lc code=end
