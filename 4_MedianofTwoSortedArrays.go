package leetcode_go

//给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n:= len(nums1), len(nums2)
	var merge []int
	i, j := 0, 0
	for ; i<m && j<n; {
		if nums1[i] <= nums2[j] {
			merge = append(merge, nums1[i])
			i++
		} else {
			merge = append(merge, nums2[j])
			j++
		}
	}
	for ; i < m; i++ {
		merge = append(merge, nums1[i])
	}
	for ; j < n; j++ {
		merge = append(merge, nums2[j])
	}
	mid := (m+n)/2
	if (m+n)%2==0 {
		return float64(merge[mid] + merge[mid-1]) / 2
	} else {
		return float64(merge[mid])
	}
}