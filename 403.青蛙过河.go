package leetcode_go

/*
 * @lc app=leetcode.cn id=403 lang=golang
 *
 * [403] 青蛙过河
 */

// @lc code=start

// 超出时间限制
// func canCross(stones []int) bool {
// 	onStone := func(cur int, next int) bool {
// 		for ; cur < len(stones); cur++ {
// 			if stones[cur] == next {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	var jump func(int, int) bool
// 	jump = func(loc int, step int) bool {
// 		if loc == len(stones)-1 {
// 			return true
// 		}
// 		if onStone(loc, stones[loc]+step) {
// 			newloc := loc
// 			for ; newloc < len(stones); newloc++ {
// 				if stones[newloc] == stones[loc]+step {
// 					break
// 				}
// 			}
// 			if newloc == loc {
// 				return false
// 			}
// 			if s1 := jump(newloc, step-1); s1 {
// 				return true
// 			}
// 			if s2 := jump(newloc, step); s2 {
// 				return true
// 			}
// 			if s3 := jump(newloc, step+1); s3 {
// 				return true
// 			}
// 			return false
// 		} else {
// 			return false
// 		}
// 	}
// 	return jump(0, 1)
// }

// 动态规划
func canCross(stones []int) bool {
	n := len(stones)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
	for i := 1; i < n; i++ {
		if stones[i]-stones[i-1] > i {
			return false
		}
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			k := stones[i] - stones[j]
			if k > j+1 {
				break
			}
			dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			if i == n-1 && dp[i][k] {
				return true
			}
		}
	}
	return false
}

// @lc code=end
// func main() {
// 	fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
// }
