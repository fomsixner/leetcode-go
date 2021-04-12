package leetcode_go

import (
	"sort"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */

// @lc code=start
func largestNumber(nums []int) string {
	var strs []string
	for i := 0; i < len(nums); i++ {
		strs = append(strs, strconv.Itoa(nums[i]))
	}
	sort.Slice(strs, func(i, j int) bool {
		temp1, _ := strconv.Atoi(strs[i] + strs[j])
		temp2, _ := strconv.Atoi(strs[j] + strs[i])
		return temp1 > temp2
	})
	res := strings.Join(strs, "")
	for len(res) > 1 && res[0]=='0' {
		res = res[1:]
	}
	return res
}

// @lc code=end


// func largestNumber(nums []int) string {
//     sort.Slice(nums, func(i, j int) bool {
//         x, y := nums[i], nums[j]
//         sx, sy := 10, 10
//         for sx <= x {
//             sx *= 10
//         }
//         for sy <= y {
//             sy *= 10
//         }
//         return sy*x+y > sx*y+x
//     })
//     if nums[0] == 0 {
//         return "0"
//     }
//     ans := []byte{}
//     for _, x := range nums {
//         ans = append(ans, strconv.Itoa(x)...)
//     }
//     return string(ans)
// }
