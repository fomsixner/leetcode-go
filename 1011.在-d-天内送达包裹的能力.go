package leetcode_go

/*
 * @lc app=leetcode.cn id=1011 lang=golang
 *
 * [1011] 在 D 天内送达包裹的能力
 */

// @lc code=start
func shipWithinDays(weights []int, D int) int {
	// 二分查找法
	// 确定边界
	left, right := weights[0], 1
	for _, v := range weights {
		right += v
		if v > left {
			left = v
		}
	}
	ans := right
	for left <= right {
		mid := (right + left) / 2
		temp, d := 0, 0
		for _, v := range weights {
			if temp+v <= mid {
				temp += v
				continue
			}
			temp = v
			d++
		}
		d++
		if d > D {
			left = mid + 1
		} else {
			if mid < ans {
				ans = mid
			}
			right = mid - 1
		}
	}
	return ans
}

// @lc code=end

// func main() {
// 	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))
// }
