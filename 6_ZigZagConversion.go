package leetcode_go

//Z字形变换

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows, res := make([]string, numRows), ""
	r, d := 0, true //d=true表示r增大
	for i, _ := range s {
		rows[r] = rows[r] + string(s[i])
		if d {
			r++
		} else {
			r--
		}
		if r == numRows {
			d = !d
			r -= 2
		}
		if r == 0 {
			d = !d
		}
	}
	for _, row := range rows {
		res = res + row
	}
	return res
}
