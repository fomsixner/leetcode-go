package leetcode_go

//给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。
//如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。
//1 <= arr.length <= 1000
//-1000 <= arr[i] <= 1000

func uniqueOccurrences(arr []int) bool {
	var m [1100]bool
	var count [2100]int
	for _, value := range arr {
		count[value+1000]++
	}
	for _, c := range count {
		if c==0 {
			continue
		}
		if m[c] == true {
			return false
		}
		if m[c] == false {
			m[c] = true
		}
	}
	return true
}

//官方解法
//首先使用哈希表记录每个数字的出现次数；随后再利用新的哈希表，统计不同的出现次数的数目。
//如果不同的出现次数的数目等于不同数字的数目，则返回 true，否则返回 false。

