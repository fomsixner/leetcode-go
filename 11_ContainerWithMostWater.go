package leetcode_go

//双指针法, 缩小问题规模
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		if height[left] < height[right] {
			if (right-left)*height[left] > res {
				res = (right-left)*height[left]
			}
			left++
		} else {
			if (right-left)*height[right] > res {
				res = (right-left)*height[right]
			}
			right--
		}
	}
	return res
}
