package leetcode_go

//动态规划
func uniquePaths(m int, n int) int {
	f := make([][]int, n+1)
	for i, _ := range f {
		f[i] = make([]int, m+1)
	}
	//初始化
	for i := 1; i <= n; i++ {
		f[i][1] = 1
	}
	for i := 1; i <= m; i++ {
		f[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 2; j <= m; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}
	return f[n][m]
}
