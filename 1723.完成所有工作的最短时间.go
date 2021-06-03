package leetcode_go

import "sort"

/*
 * @lc app=leetcode.cn id=1723 lang=golang
 *
 * [1723] 完成所有工作的最短时间
 */

// @lc code=start
func minimumTimeRequired(jobs []int, k int) int {
	total := 0
	maxTime := 0
	for _, v := range jobs {
		if v > maxTime {
			maxTime = v
		}
		total += v
	}
	avg := total / k
	sort.Ints(jobs)
	left, right := 0, len(jobs)
	for left < right {
		temp := 0
		
	}

}
// @lc code=end

