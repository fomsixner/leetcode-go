package leetcode_go

// 求子集
// 迭代法
// 对于原集合中的每一个元素，只有在子集中和不在子集中这2个状态，用1表示在子集中，0表示不在子集中
// 可以得到一个二进制0,1序列
func subsets(nums []int) [][]int {
	var res [][]int
	for mask := 0; mask < (1 << len(nums)); mask++ {
		var subset []int
		for i, v := range nums {
			if ((mask >> i) & 1) > 0 {
				subset = append(subset, v)
			}
		}
		res = append(res, subset)
	}
	return res
}
