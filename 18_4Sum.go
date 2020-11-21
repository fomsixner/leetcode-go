package leetcode_go

import "sort"

//给定一个包含n 个整数的数组nums和一个目标值target，判断nums中是否存在四个元素 a，b，c和 d，
//使得a + b + c + d的值与target相等？ 找出所有满足条件且不重复的四元组。

//双指针
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for a := 0; a < len(nums)-3; a++ {
		if a != 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < len(nums)-2; b++ {
			if b != a+1 && nums[b] == nums[b-1] {
				continue
			}
			for c, d := b+1, len(nums)-1; c < d; {
				if c != b+1 && nums[c] == nums[c-1] {
					c++
					continue
				}
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum == target {
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					c++
				} else if sum < target {
					c++
				} else {
					d--
				}
			}
		}
	}
	return res
}
