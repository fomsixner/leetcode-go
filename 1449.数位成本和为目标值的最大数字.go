package leetcode_go

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=1449 lang=golang
 *
 * [1449] 数位成本和为目标值的最大数字
 */

// @lc code=start

// 动态规划
// dp[i][j] 表示用前i个数凑成targrt=j时得到的最大整数
// 此版本时间复杂度较高，内存开销较大
func largestNumber1449(cost []int, target int) string {
	max := func(a, b string) string {
		if len(a) > len(b) {
			return a
		} else if len(a) < len(b) {
			return b
		}
		if strings.Compare(a, b) > 0 {
			return a
		} else {
			return b
		}
	}
	combination := func(a string, b string) string {
		// 用拼接再排序的方法超时
		// res := strings.Split(a+b, "")
		// sort.Sort(sort.Reverse(sort.StringSlice(res)))
		// return strings.Join(res, "")
		// b 是有序的字符串，因此只要用插入排序遍历一遍就行了
		res := []byte(b)
		add := []byte(a)[0]
		i := 0
		for ; i < len(res); i++ {
			if res[i] < add {
				break
			}
		}
		res = append(res[0:i], append([]byte{add}, res[i:]...)...)
		return string(res)
	}

	dp := make([]string, target+1)
	for i, cost := range cost {
		a := strconv.Itoa(i + 1)
		for j := cost; j <= target; j++ {
			if dp[j-cost] != "" || j-cost == 0 {
				dp[j] = max(dp[j], combination(a, dp[j-cost]))
			}

		}
	}
	if dp[target] == "" {
		return "0"
	}
	return dp[target]
}

// @lc code=end

// func main() {
// 	fmt.Println(largestNumber1449([]int{4,3,2,5,6,7,2,5,5}, 9))
// }
