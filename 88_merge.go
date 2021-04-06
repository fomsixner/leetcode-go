package leetcode_go

import "sort"

// 排序法
func merge1(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 双指针法
func merge2(nums1 []int, m int, nums2 []int, n int) {
	ans := []int{}
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			ans = append(ans, nums1[i])
			i++
		} else {
			ans = append(ans, nums2[j])
			j++
		}
	}
	if i != m {
		ans = append(ans, nums1[i:]...)
	}
	if j != n {
		ans = append(ans, nums2[j:]...)
	}
	copy(nums1, ans)
}

// 逆向双指针法
// 指针设置为从后向前遍历，每次取两者之中的较大者放进nums1的最后面
func merge3(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	cur := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[cur] = nums1[i]
			i--
		} else {
			nums1[cur] = nums2[j]
			j--
		}
		cur--
	}
	if j >= 0 {
		copy(nums1[0:], nums2[0:j+1])
	}
}
