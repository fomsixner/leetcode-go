package leetcode_go

/*
 * @lc app=leetcode.cn id=65 lang=golang
 *
 * [65] 有效数字
 */

// @lc code=start
func isNumber(s string) bool {
	count_point, count_e := 0, 0
	e_position, point_position := -1, -1
	for i, v := range s {
		if v == '.' {
			count_point++
			point_position = i
		} else if v == 'e' || v == 'E' {
			count_e++
			e_position = i
		} else if v != '+' && v != '-' && (v < '0' || v > '9') {
			return false
		}
	}
	if count_point > 1 || count_e > 1 || (e_position != -1 && point_position > e_position) {
		return false
	}
	if e_position == len(s)-1 {
		return false
	}
	var left, right string
	if e_position == -1 {
		left, right = s, ""
	} else {
		left, right = s[0:e_position], s[e_position+1:]
	}
	flag := false
	for i, v := range left {
		if i != 0 && (v == '+' || v == '-') {
			return false
		}
		if v >= '0' && v <= '9' {
			flag = true
		}
	}
	if !flag {
		return false
	}
	flag = false
	for i, v := range right {
		if i != 0 && (v == '+' || v == '-') {
			return false
		}
		if v >= '0' && v <= '9' {
			flag = true
		}
	}
	if len(right) != 0 && !flag {
		return false
	}
	return true
}

// 确定有穷状态自动机
// func isNumber(s string) bool {
// 	type charType int
// 	type state int
// 	const (
// 		start state = iota
// 		sign
// 		point
// 		integer
// 		point_without_int
// 		decimal
// 		exp
// 		exp_sign
// 		exp_integer
// 		end
// 	)
// 	const (
// 		char_number charType = iota
// 		char_point
// 		char_sign
// 		char_exp
// 		char_illegal
// 	)

// 	switchType := func(ch byte) charType {
// 		switch ch {
// 		case '0','1','2','3','4','5','6','7','8','9':
// 			return char_number
// 		case '+','-':
// 			return char_sign
// 		case 'e', 'E':
// 			return char_exp
// 		case '.':
// 			return char_point
// 		default:
// 			return char_illegal
// 		}
// 	}
// 	current := start
// 	for i, _ := range s {
// 		switch current {
// 		case start:
// 			ch_type := switchType(s[i])
// 			if ch_type == char_sign {
// 				current = sign
// 			} else if ch_type == char_number {
// 				current = integer
// 			} else if ch_type == char_point {
// 				current = point
// 			} else {
// 				return false
// 			}
// 		}
// 	}

// }

// @lc code=end
