package leetcode_go

//每次消去的数num[i]为 num[i]>=num[i-1] && num[i]>num[i+1]

func removeKdigits(num string, k int) string {
	for i := 0; i < k; i++ {
		j := 0
		for ; j < len(num)-1 && num[j] <= num[j+1]; j++ {
		}
		num = num[:j] + num[j+1:]
	}
	//去除首部的0
	for len(num) > 0 && num[0] == '0' {
		num = num[1:]
	}
	//如果长度为0则返回0
	if len(num) == 0 {
		return "0"
	}
	return num
}
