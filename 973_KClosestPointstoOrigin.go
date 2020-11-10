package leetcode_go

import (
	"sort"
)

//方法一: 排序
func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {        //不需要计算平方
		return float64(points[i][0]*points[i][0] + points[i][1]*points[i][1]) <
			float64(points[j][0]*points[j][0] + points[j][1]*points[j][1])
	})
	//var res [][]int
	//for i:=0; i<K; i++ {
	//	res = append(res, points[i])
	//}
	//return res
	return points[:K]
}


//方法二: 优先队列
