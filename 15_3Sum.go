package leetcode_go

import "sort"

//双指针法
func threeSum(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	sort.Ints(nums) //从小到大排序
	for i := 0; i <= n-3; i++ {
		if i == 0 || nums[i] != nums[i-1] { //排除重复
			for j, k := i+1, n-1; j < k; {
				if j == i+1 || nums[j] != nums[j-1] { //排除重复
					sum := nums[i] + nums[j] + nums[k]
					if sum == 0 {
						res = append(res, []int{nums[i], nums[j], nums[k]})
						j++
					} else if sum < 0 {
						j++
					} else {
						k--
					}
				} else {
					j++
				}
			}
		}
	}
	return res
}
