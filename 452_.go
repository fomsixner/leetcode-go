package leetcode_go

import "sort"

//找重叠区间
func findMinArrowShots(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		} else {
			return points[i][1] < points[j][1]
		}
	})
	res := 1
	left, right := points[0][0], points[0][1]
	for _, p := range points {
		if p[0] >= left && p[0] <= right {
			left = p[0]
			if p[1] < right {
				right = p[1]
			}
		} else {
			res++
			left, right = p[0], p[1]
		}
	}
	return res
}
