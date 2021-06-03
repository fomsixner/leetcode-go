package leetcode_go

/*
 * @lc app=leetcode.cn id=740 lang=golang
 *
 * [740] 删除与获得点数
 */

// @lc code=start
func deleteAndEarn(nums []int) int {
	maxNum := 0
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	
	mp := make([]int, maxNum+1)
	for _, v := range nums {
		mp[v] += v
	}

	first, second := mp[0], max(mp[0], mp[1])
	for i := 2; i < len(mp); i++ {
		first, second = second, max(first+mp[i], second)
	}
	return second
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end
