package leetcode_go

//1和0
//动态规划, 0-1背包问题

func findMaxForm(strs []string, m int, n int) int {
	length := len(strs)
	zeros := make([]int, length)
	ones := make([]int, length)
	for i, str := range strs { //获得每个字符串的0,1个数
		for _, c := range str {
			if c == '1' {
				ones[i] += 1
			} else {
				zeros[i] += 1
			}
		}
	}
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	for k := 0; k < length; k += 1 {
		for i := m; i >= zeros[k]; i -= 1 {
			for j := n; j >= ones[k]; j -= 1 {
				if (dp[i-zeros[k]][j-ones[k]] + 1) > dp[i][j] {
					dp[i][j] = dp[i-zeros[k]][j-ones[k]] + 1
				}
			}
		}
	}
	return dp[m][n]
}
