package leetcode_go

import "sort"

//哈希表法
func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int, 10)
	for _, v := range arr2 {
		m[v] = 1
	}
	var tail, res []int
	for _, v := range arr1 {
		if m[v] != 0 {
			m[v]++
		} else {
			tail = append(tail, v)
		}
	}
	sort.Ints(tail)
	for _, v := range arr2 {
		for i := 1; i < m[v]; i++ {
			res = append(res, v)
		}
	}
	for _, v := range tail {
		res = append(res, v)
	}
	return res
}
