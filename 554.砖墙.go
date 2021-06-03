package leetcode_go

/*
 * @lc app=leetcode.cn id=554 lang=golang
 *
 * [554] 砖墙
 */

// 哈希表法
// 穿过的砖块越少，穿过的边缘越多

// @lc code=start
func leastBricks(wall [][]int) int {
	ans := len(wall)
	mp := make(map[int]int)
	for _, row := range wall {
		distance := 0
		for i := 0; i < len(row)-1; i++ {
			mp[row[i]+distance]++
			distance += row[i]
		}
	}
	edge := 0
	for _, v := range mp {
		if v > edge {
			edge = v
		}
	}
	return ans - edge
}

// @lc code=end
