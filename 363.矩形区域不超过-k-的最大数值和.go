package leetcode_go

/*
 * @lc app=leetcode.cn id=363 lang=golang
 *
 * [363] 矩形区域不超过 K 的最大数值和
 */

// @lc code=start
func maxSumSubmatrix(matrix [][]int, k int) int {
	n, m := len(matrix), len(matrix[0])
	// 创建前缀
	S := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		S[i] = make([]int, m+1)
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			S[i][j] = S[i-1][j] + S[i][j-1] - S[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	// 暴力搜索
	ans := ^int(^uint(0) >> 1) // 获取int最小值
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for p := i + 1; p < n+1; p++ {
				for q := j + 1; q < m+1; q++ {
					temp := S[p][q] - S[i][q] - S[p][j] + S[i][j]
					if temp <= k && temp > ans {
						ans = temp
					}
				}
			}
		}
	}
	return ans
}

// @lc code=end

// func main(){
// 	data := []int{2,2,-1}
// 	fmt.Println(maxSumSubmatrix([][]int{data}, 0))
// }
