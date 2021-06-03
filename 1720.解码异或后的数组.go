package leetcode_go

/*
 * @lc app=leetcode.cn id=1720 lang=golang
 *
 * [1720] 解码异或后的数组
 */

// A xor B xor B = A

// @lc code=start
func decode(encoded []int, first int) []int {
	ans := []int{first}
	for _, v := range encoded {
		ans = append(ans, ans[len(ans)-1]^v)
	}
	return ans
}

// @lc code=end
