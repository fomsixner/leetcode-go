package leetcode_go

import "sort"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		if (nums[i+1] - nums[i]) > res {
			res = nums[i+1] - nums[i]
		}
	}
	return res
}
