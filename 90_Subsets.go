package leetcode_go

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for mask := 0; mask < (1 << len(nums)); mask++ {
		var subset []int
		flag := 1
		for i, v := range nums {
			if (mask >> i & 1) > 0 {
				// 如果与前一个数相同且没有选择前一个数,则可跳过当前循环
				if i > 0 && mask>>(i-1)&1 == 0 && v == nums[i-1] {
					flag = 0
					break
				}
				subset = append(subset, v)
			}
		}
		if flag == 1 {
			res = append(res, subset)
		}
	}
	return res
}
