package leetcode_go

/*
 * @lc app=leetcode.cn id=1442 lang=golang
 *
 * [1442] 形成两个异或相等数组的三元组数目
 */

// @lc code=start

// 暴力：超时
// func countTriplets(arr []int) int {
// 	ans := 0
// 	for i := 0; i < len(arr); i++ {
// 		for j := i + 1; j < len(arr); j++ {
// 			for k := j; k < len(arr); k++ {
// 				a, b := 0, 0
// 				for x := i; x <= j-1; x++ {
// 					a ^= arr[x]
// 				}
// 				for y := j; y <= k; y++ {
// 					b ^= arr[y]
// 				}
// 				if a == b {
// 					ans++
// 				}
// 			}
// 		}
// 	}
// 	return ans
// }

// 使用前缀和+二重循环
// func countTriplets(arr []int) int {
// 	pre := make([]int, len(arr)+1)
// 	pre[0] = 0
// 	for i := 0; i < len(arr); i++ {
// 		pre[i+1] = pre[i] ^ arr[i]
// 	}
// 	ans := 0
// 	for i := 0; i < len(arr); i++ {
// 		// for j := i + 1; j < len(arr); j++ {
// 		// a := pre[i] ^ pre[j]
// 		for k := i + 1; k < len(arr); k++ { // 可以缩短到二重循环
// 			// b := pre[j] ^ pre[k+1]
// 			// if a == b {
// 			// 	ans++
// 			// }
// 			if pre[i] == pre[k+1] { // 与j无关
// 				// ans++
// 				ans += k - i
// 			}
// 		}
// 		// }
// 	}
// 	return ans
// }

// 使用哈希表，一重循环
func countTriplets(arr []int) int {
	ans := 0
	n := len(arr)
	s := make([]int, n+1)
	for i, v := range arr {
		s[i+1] = s[i] ^ v
	}
	cnt := map[int]int{}
	total := map[int]int{}
	for k := 0; k < n; k++ {
		if m, has := cnt[s[k+1]]; has {
			ans += m*k - total[s[k+1]]
		}
		cnt[s[k]]++
		total[s[k]] += k
	}
	return ans
}

// @lc code=end
