package leetcode_go
/*
 * @lc app=leetcode.cn id=1310 lang=golang
 *
 * [1310] 子数组异或查询
 */

// @lc code=start


// 执行速度太慢
// func xorQueries(arr []int, queries [][]int) []int {
// 	ans := []int{}
// 	for _, query := range queries {
// 		temp := 0
// 		for i:= query[0]; i<=query[1]; i++ {
// 			temp = temp ^ arr[i]
// 		}
// 		ans = append(ans, temp)
// 	}
// 	return ans
// }


//前缀异或
func xorQueries(arr []int, queries [][]int) []int { 
	xors := make([]int, len(arr)+1)
	for i, v := range arr {
		xors[i+1] = xors[i] ^ v
	}
	ans := []int{}
	for _, query := range queries {
		ans = append(ans, xors[query[0]] ^ xors[query[1]+1])
	}
	return ans
}


// @lc code=end

