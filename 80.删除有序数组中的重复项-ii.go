package leetcode_go

// @lc code=start
func removeDuplicates(nums []int) int {
	count := 1
	n := len(nums)
	for i := 1; i < n; i++ {
		for i < n && nums[i] == nums[i-1] {
			i++
			count++
		}
		if count > 2 {
			newloc := i - count + 2
			for j := i; j < n; j++ {
				nums[newloc] = nums[j]
				newloc++
			}
			i = i - count + 2
			n = n - count + 2
		}
		count = 1
	}
	return n
}

// @lc code=end


// 双指针法
// func removeDuplicates(nums []int) int {
//     n := len(nums)
//     if n <= 2 {
//         return n
//     }
//     slow, fast := 2, 2
//     for fast < n {
//         if nums[slow-2] != nums[fast] {
//             nums[slow] = nums[fast]
//             slow++
//         }
//         fast++
//     }
//     return slow
// }
