package leetcode_go

/*
 * @lc app=leetcode.cn id=525 lang=golang
 *
 * [525] 连续数组
 */

// @lc code=start

// 使用前缀和+哈希表 
func findMaxLength(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	mp := make(map[int]int)
	mp[0] = -1
	counter := 0
	ans := 0
	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}
		if value, ok := mp[counter]; ok { // value为第一次出现该前缀值的下标
			ans = max(ans, i-value)
		} else {
			mp[counter] = i
		}
	}
	return ans
}

// @lc code=end
