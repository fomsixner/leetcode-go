package leetcode_go

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		j := 0
		for ; j < len(needle); j++ {
			if (i+j < len(haystack) && haystack[i+j] != needle[j]) || i+j >= len(haystack) {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}

// @lc code=end
