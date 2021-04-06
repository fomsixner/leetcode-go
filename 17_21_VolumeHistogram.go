package leetcode_go

// 动态规划法
func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	ans := 0
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	for i := 0; i < n; i++ {
		ans += min(rightMax[i], leftMax[i]) - height[i]
	}
	return ans
}

// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }

// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }
