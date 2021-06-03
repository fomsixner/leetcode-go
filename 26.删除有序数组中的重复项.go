package leetcode_go

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates1(nums []int) int {
	left, cur := 0, 1
	for ; cur < len(nums); cur++ {
		if nums[left] == nums[cur] {
			continue
		}
		nums[left+1] = nums[cur]
		left++
	}
	return left + 1
}

// @lc code=end
