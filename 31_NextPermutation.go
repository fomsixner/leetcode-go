package leetcode_go

//两遍扫描法

func reverse(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

func nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	i, j := n-2, n-1
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}
	if i == -1 {
		reverse(nums)
	} else {
		for ; j >= i+1 && nums[j] <= nums[i]; j-- {
		}
		nums[i], nums[j] = nums[j], nums[i]
		reverse(nums[i+1:])
	}
}

//func main()  {
//	nums := []int{5, 1, 1}
//	nextPermutation(nums)
//}
