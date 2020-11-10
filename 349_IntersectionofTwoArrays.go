package leetcode_go

//两个数组的交集


func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, s:= range nums1 {
		m[s] = 1
	}
	res := []int{}
	for _, s:= range nums2 {
		if m[s] == 1 {
			m[s] += 1       //防止重复添加
			res = append(res, s)
		}
	}
	return res
}
