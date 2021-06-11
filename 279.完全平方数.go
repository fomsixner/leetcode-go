package leetcode_go

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// 动态规划, dp[i][j] 表示从数1~i的平方组合成j所需的最小数量

// @lc code=start
// 初始版本
// func numSquares(n int) int {
// 	m := int(math.Sqrt(float64(n)))
// 	dp := make([][]int, m+1)
// 	for i := range dp {
// 		dp[i] = make([]int, n+1)
// 	}
// 	for i := 0; i <= n; i++ {
// 		dp[1][i] = i
// 	}
// 	for i := 2; i <= m; i++ {
// 		for j := 0; j <= n; j++ {
// 			dp[i][j] = dp[i-1][j]
// 			for k := 0; k*i*i <= j; k++ {
// 				if dp[i-1][j-k*i*i]+k < dp[i][j] {
// 					dp[i][j] = dp[i][j-k*i*i] + k
// 				}
// 			}
// 		}
// 	}
// 	return dp[m][n]
// }

// 优化内存
// func numSquares(n int) int {
// 	m := int(math.Sqrt(float64(n)))
// 	dp := make([]int, n+1)
// 	for i := 0; i <= n; i++ {
// 		dp[i] = i
// 	}
// 	for i := 2; i <= m; i++ {
// 		for j := 0; j <= n; j++ {
// 			for k := 0; k*i*i <= j; k++ {
// 				if dp[j-k*i*i]+k < dp[j] {
// 					dp[j] = dp[j-k*i*i] + k
// 				}
// 			}
// 		}
// 	}
// 	return dp[n]
// }

// 时间复杂度优化
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = i
	}
	for i := 2; i*i <= n; i++ {
		k := i * i
		for j := k; j <= n; j++ {
			if dp[j-k]+1 < dp[j] {
				dp[j] = dp[j-k] + 1
			}
		}
	}
	return dp[n]
}

// @lc code=end

// func main() {
// 	fmt.Println(numSquares(13))
// }
