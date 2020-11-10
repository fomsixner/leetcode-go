package leetcode_go


//func isMatch(s string, p string) bool {
//	i, j := 0, 0
//	any := false
//	n, m := len(s), len(p)
//	for n > 0 && m > 0 && s[n-1] == p[m-1] {
//		n, m = n-1, m-1
//	}
//	for i < n {
//		if j >= m {return false}
//		if p[j] == s[i] {
//			i, j = i+1, j+1
//			any = false
//		} else if p[j] == '.' {
//			i, j = i+1, j+1
//			any = true
//			continue
//		} else if p[j] == '*' {
//			if any {
//				if j != m-1 {return false}
//				return true
//			}
//			temp := p[j-1]
//			for s[i] == temp {
//				i++
//				if i == n {
//					break
//				}
//			}
//			j++
//		} else {
//			if p[j+1] != '*' {return false}
//			j++
//		}
//	}
//	if j < m {return false}
//	return true
//}
//

//动态规划
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m + 1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n + 1)
	}
	f[0][0] = true         //表示空字符串的匹配
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j - 1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}


//func main()  {
//	s := "aaa"
//	p := "ab*ac*a"
//	if isMatch(s, p) {
//		fmt.Println("true")
//	} else {
//		fmt.Print("false")
//	}
//}
