package leetcode_go

import "math"

//动态规划
//每一个字符在ring中可以有多个位置匹配, 动态规划相当于把每种可能都尝试了一遍
//采用dp[i][j]表示用ring中位置j的字符来匹配key中位置i的字符

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findRotateSteps(ring string, key string) int {
	n, m := len(ring), len(key)
	var pos [26][]int
	for i, v := range ring {
		pos[v-'a'] = append(pos[v-'a'], i)
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt64
		}
	}
	for _, v := range pos[key[0]-'a'] {
		dp[0][v] = min(n-v, v) + 1
	}
	for i := 1; i < m; i++ {
		for _, j := range pos[key[i]-'a'] {
			for _, k := range pos[key[i-1]-'a'] {
				dp[i][j] = min(dp[i][j], dp[i-1][k]+min(abs(j-k), n-abs(j-k))+1)
			}
		}
	}
	steps := math.MaxInt64
	for _, v := range pos[key[m-1]-'a'] {
		if dp[m-1][v] < steps {
			steps = dp[m-1][v]
		}
	}
	return steps
}
