package leetcode_go

/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	ans := ""
	for _, v := range num2 {
		carry := 0
		b := int(v - '0')
		temp := make([]byte, len(num1)+1)
		for i := len(num1) - 1; i >= 0; i-- {
			a := int(num1[i] - '0')
			cur := (carry + a*b) % 10
			carry = (carry + a*b) / 10
			temp[i+1] = byte(cur + '0')
		}
		temp[0] = byte(carry + '0')
		if ans == "" {
			ans = string(temp)
		} else {
			res := []byte(ans)
			res = append(res, '0')
			carry = 0
			i, j := len(temp)-1, len(res)-1
			for i >= 0 && j >= 0 {
				a := int(res[j] - '0')
				b := int(temp[i] - '0')
				cur := (carry + a + b) % 10
				carry = (carry + a + b) / 10
				res[j] = byte(cur + '0')
				i--
				j--
			}
			for j >= 0 && carry > 0 {
				a := int(res[j] - '0')
				cur := (carry + a) % 10
				carry = (carry + a) / 10
				res[j] = byte(cur + '0')
				j--
			}
			ans = string(res)
		}
	}
	for ans[0] == '0' && len(ans) > 1 {
		ans = ans[1:]
	}
	return ans
}

// @lc code=end
