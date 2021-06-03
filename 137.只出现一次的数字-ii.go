package leetcode_go

/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
// hash表法
// func singleNumber(nums []int) int {
// 	mp := make(map[int]int)
// 	for _, v:= range nums {
// 		mp[v]++
// 	}
// 	for k, v := range mp {
// 		if v == 1 {
// 			return k
// 		}
// 	}
// 	return 0    // 不会发生
// }

// 位运算法
// 有限状态自动机
// 设计一个转换电路，使一个数出现3次时自动抵消为0
// sum % 3 的结果为0,1,2 共3个状态
// 记录3个状态需要2个二进制位
func singleNumber(nums []int) int {
	one, two := 0, 0
	for _, v := range nums {
		one = one ^ v & ^two
		two = two ^ v & ^one
	}
	return one
}

// @lc code=end
