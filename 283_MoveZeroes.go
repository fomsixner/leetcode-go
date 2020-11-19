package leetcode_go

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
func moveZeroes(nums []int) {
	n, j := len(nums), 1
	for i := 0; i < n; {
		for ; i < n && nums[i] != 0; i++ {
		}
		for j = i + 1; j < n && nums[j] == 0; j++ {
		}
		if j >= n {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
}

//方法二
//使用双指针，左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部。
//右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。
//相当于每次找到非0的数移动到已处理好的末尾
func moveZeroes2(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
