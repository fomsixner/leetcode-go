package leetcode_go

import "sort"

//一般这种数对，还涉及排序的，根据第一个元素正向排序，根据第二个元素反向排序，
//或者根据第一个元素反向排序，根据第二个元素正向排序，往往能够简化解题过程

//在本题目中，我首先对数对进行排序，按照数对的元素 1 降序排序，按照数对的元素 2 升序排序。
//原因是，按照元素 1 进行降序排序，对于每个元素，在其之前的元素的个数，就是大于等于他的元素的数量，
//而按照第二个元素正向排序，我们希望 k 大的尽量在后面，减少插入操作的次数。

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] != people[j][0] {
			return people[i][0] > people[j][0]
		} else {
			return people[i][1] < people[j][1]
		}
	})
	var res [][]int
	for i, _ := range people {
		if people[i][1] >= len(res) {
			res = append(res, people[i])
		} else {
			res = append(res[:people[i][1]], append([][]int{people[i]}, res[people[i][1]:]...)...)
		}
	}
	return res
}
