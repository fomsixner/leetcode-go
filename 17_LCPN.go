package leetcode_go

//回溯
func letterCombinations(digits string) []string {
	mp := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var res []string
	if len(digits) == 0 {
		return res
	}
	var search func(str string, temp string)
	search = func(str string, temp string) {
		for _, v := range mp[str[0]-'0'] {
			cur := temp + string(v)
			if len(str) > 1 {
				search(str[1:], cur)
			} else {
				res = append(res, cur)
			}
		}
	}
	search(digits, "")
	return res
}

//队列
//先将输入的 digits 中第一个数字对应的每一个字母入队，然后将出队的元素与第二个数字对应的每一个字母组合后入队...
//直到遍历到 digits 的结尾。最后队列中的元素就是所求结果
//

func letterCombinations2(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mp := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	queue := []string{""}
	for _, v := range digits {
		var n int = len(queue)
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			for _, c := range mp[v-'0'] {
				queue = append(queue, cur+string(c))

			}
		}
	}
	return queue
}
