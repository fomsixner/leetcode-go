package leetcode_go

import "sort"

/*
 * @lc app=leetcode.cn id=1738 lang=golang
 *
 * [1738] 找出第 K 大的异或坐标值
 */

// @lc code=start

// 使用前缀和
func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	pre := make([][]int, m+1)
	for i, _ := range pre {
		pre[i] = make([]int, n+1)
	}
	list := []int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pre[i+1][j+1] = pre[i][j+1] ^ pre[i+1][j] ^ pre[i][j] ^ matrix[i][j]
			list = append(list, pre[i+1][j+1])
		}
	}
	sort.Ints(list)
	return list[len(list)-k]
}

// @lc code=end
