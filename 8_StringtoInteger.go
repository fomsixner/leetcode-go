package leetcode_go

import (
	"math"
)

//请你来实现一个 atoi 函数，使其能将字符串转换成整数。
//
//首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。接下来的转化规则如下：
//
//如果第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字字符组合起来，形成一个有符号整数。
//假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成一个整数。
//该字符串在有效的整数部分之后也可能会存在多余的字符，那么这些字符可以被忽略，它们对函数不应该造成影响。
//注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换，即无法进行有效转换。
//
//在任何情况下，若函数不能进行有效的转换时，请返回 0 。


func myAtoi(s string) int {
	i, n:= 0, len(s)
	for ; i < n && s[i] == ' '; i++ {}           //跳过空格
	if i == n || ((s[i] != '-' && s[i] != '+') && (s[i] < '0' || s[i] > '9')) {
		return 0
	}
	sign, res := 1, 0
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}
	for ; i < n; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if res > (math.MaxInt32 - int(s[i] -'0')) / 10 {  //溢出
				if sign == 1 {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}
			res = res*10 + int(s[i] -'0')
		} else {
			break
		}
	}
	return res*sign
}

//func main()  {
//	s := "-91283472332"
//	fmt.Println(myAtoi(s))
//}