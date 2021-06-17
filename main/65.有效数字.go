package main

import "fmt"

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

// @lc code=end

func main() {
	fmt.Println(isNumber("e.7e5"))
}
