package leetcode_go

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxLength, wiggle := 1, 0
	pre, cur := 0, 1
	for cur < len(nums) {
		if nums[cur] < nums[pre] && wiggle != -1 {
			maxLength++
			wiggle = -1
			pre = cur
			cur++
		} else if nums[cur] > nums[pre] && wiggle != 1 {
			maxLength++
			wiggle = 1
			pre = cur
			cur++
		} else {
			temp := cur
			for temp+1 < len(nums) && nums[temp+1]*wiggle > nums[temp]*wiggle {
				temp++
			}
			pre = temp
			cur = temp + 1
		}
	}
	return maxLength
}

//func main() {
//	nums := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
//	fmt.Println(wiggleMaxLength(nums))
//}
