package leetcode_go

import "strings"

//每次消去的数num[i]为 num[i]>=num[i-1] && num[i]>num[i+1]

func removeKdigits(num string, k int) string {
	for i := 0; i < k; i++ {
		j := 0
		for ; j < len(num)-1 && num[j] <= num[j+1]; j++ {
			//这里可以优化一下,每次不是从头开始遍历,而是从位置j-1开始,如果比num[j]大则消去num[j-1],并把j-1
		}
		num = num[:j] + num[j+1:]
	}
	//去除首部的0
	num = strings.TrimLeft(num, "0")
	//如果长度为0则返回0
	if len(num) == 0 {
		return "0"
	}
	return num
}

func removeKdigits2(num string, k int) string {
	j := 0
	for i := 0; i < k; i++ {
		for ; j < len(num)-1 && num[j] <= num[j+1]; j++ {
		}
		num = num[:j] + num[j+1:]
		if j > 0 {
			j--
		}
	}
	//去除首部的0
	num = strings.TrimLeft(num, "0")
	//如果长度为0则返回0
	if len(num) == 0 {
		return "0"
	}
	return num
}

//官方解法,单调栈
//考虑从左往右增量的构造最后的答案。我们可以用一个栈维护当前的答案序列，栈中的元素代表截止到当前位置，
//删除不超过 kk 次个数字后，所能得到的最小整数。根据之前的讨论：在使用 kk 个删除次数之前，栈中的序列从栈底到栈顶单调不降。
//
//因此，对于每个数字，如果该数字小于栈顶元素，我们就不断地弹出栈顶元素，直到
//
//栈为空
//或者新的栈顶元素不大于当前数字
//或者我们已经删除了 kk 位数字

func removeKdigits3(num string, k int) string {
	stack := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}
