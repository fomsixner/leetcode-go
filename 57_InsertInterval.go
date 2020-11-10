package leetcode_go

//给出一个无重叠的 ，按照区间起始端点排序的区间列表。
//
//在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	left, right:= newInterval[0], newInterval[1]
	inserted := false
	for _, interval:= range intervals{
		if left > interval[1] {   //无交集
			res = append(res, interval)
		} else if right < interval[0] {  //无交集
			if !inserted {
				res = append(res, []int{left, right})
				inserted = true
			}
			res = append(res, interval)
		} else {
			if left > interval[0] {
				left = interval[0]
			}
			if right < interval[1] {
				right = interval[1]
			}
		}
	}
	if !inserted {   //插到最后
		res = append(res, []int{left, right})
	}
	return res
}
