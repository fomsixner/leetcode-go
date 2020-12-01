package leetcode_go

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	mid := (left + right) / 2
	for left < right && nums[mid] != target {
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left + right) / 2
	}
	if nums[mid] != target {
		return []int{-1, -1}
	}
	for left = mid; left-1 >= 0 && nums[left-1] == target; left-- {
	}
	for right = mid; right+1 < len(nums) && nums[right+1] == target; right++ {
	}
	return []int{left, right}
}
