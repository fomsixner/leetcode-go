package leetcode_go

//单调栈
//假设从nums1中取出x个数字
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	x := 1
	if len(nums2)-k < 0 {
		x = k - len(nums2)
	}
	var res []int
	if k < len(nums1) {
		res = nums1[:k]
	} else {
		res = append(nums1, nums2[:k-len(nums1)]...)
	}
	for x <= k && x <= len(nums1) && (k-x) <= len(nums2) {
		arr1, arr2 := findMaxArray(x, nums1), findMaxArray(k-x, nums2)
		temp := mergeArr(arr1, arr2)
		for i := 0; i < k; i++ {
			if temp[i] > res[i] {
				res = temp
				break
			} else if temp[i] < res[i] {
				break
			}
		}
		x++
	}
	return res
}

func findMaxArray(x int, nums []int) []int {
	if x == 0 {
		return []int{}
	}
	var stack []int
	stack = append(stack, nums[0])
	count := 1
	for i := 1; i < len(nums); i++ {
		for count > 0 && nums[i] > stack[len(stack)-1] && i+x-count < len(nums) {
			stack = stack[:len(stack)-1]
			count--
		}
		if count < x {
			stack = append(stack, nums[i])
			count++
		}
	}
	return stack
}

func mergeArr(nums1, nums2 []int) []int {
	var res []int
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if !lexicographicalLess(nums1[i:], nums2[j:]) {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	if i < len(nums1) {
		res = append(res, nums1[i:]...)
	}
	if j < len(nums2) {
		res = append(res, nums2[j:]...)
	}
	return res
}

func lexicographicalLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

//func main()  {
//	nums1 := []int{8,5,9,5,1,6,9}
//	nums2 := []int{2,6,4,3,8,4,1,0,7,2,9,2,8}
//	k := 20
//	fmt.Println(maxNumber(nums1, nums2, k))
//}
