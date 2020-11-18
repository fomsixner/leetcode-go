package leetcode_go

import (
	"math"
	"sort"
)

//方法一: 计算每个点的曼哈顿距离,排序
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	var res [][]int
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			res = append(res, []int{i, j})
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return math.Abs(float64(res[i][0]-r0))+math.Abs(float64(res[i][1]-c0)) <
			math.Abs(float64(res[j][0]-r0))+math.Abs(float64(res[j][1]-c0))
	})
	return res
}

//方法二: 桶排序
//遍历所有坐标，按照距离的大小分组，每组的距离相等（即放入一个桶中）
//按照距离从小到大的原则，遍历所有桶，并输出结果

//方法三: BFS
//可以把所有的坐标看作树的结点，距离相等的结点位于树的同一层
//而对于每一层的结点，它们的距离 dist 可以分为行距离和列距离，且 rowDist + colDist = dist 必然成立
//使 rowDist 从 0 到 dist 递增，colDist 相应有不同的值，可以得到不同的坐标：
//横坐标为：r0 - rowDist 或 r0 + rowDist
//纵坐标为：c0 - colDist 或 c0 + colDist
//注意特殊情况：rowDist 或 colDist 为 0 时，每组只有一个正确值
//对步骤 3 中，所有在矩阵范围内的坐标进行记录
