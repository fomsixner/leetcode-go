package leetcode_go

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// 双指针法
// @lc code=start
func removeElement(nums []int, val int) int {
	left, cur := 0, 0
	for ; cur < len(nums); cur++ {
		if nums[cur] == val {
			continue
		}
		nums[left] = nums[cur]
		left++
	}
	return left
}

// @lc code=end
