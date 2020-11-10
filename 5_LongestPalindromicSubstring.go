package leetcode_go

//最长回文子串

//方法1: 暴力破解, 列举所有子串
func isPalindromeStr(str string) bool {
	n := len(str)
	for i:=0; i < n/2; i+=1 {
		if str[i] != str[n-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	maxlength := 0
	var maxPalindrome string
	n := len(s)
	if n == 1 {
		return s
	}
	for i:=0; i < n-1; i += 1 {
		for j:=i+1; j <= n; j+=1 {
			if isPalindromeStr(s[i:j]) && j-i+1 > maxlength {
				maxlength = j-i+1
				maxPalindrome = s[i:j]
			}
		}
	}
	return maxPalindrome
}

//方法2: 动态规划
func longestPalindromeDP(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i, _ := range dp {
		dp[i] = make([]bool, n)
	}
	//初始化, 单个字符必定为回文字符串
	for i:=0; i < n; i+=1 {
		dp[i][i] = true
	}
	maxlength := 0
	start, end := 0, 0       //[start, end]
	for j:=1; j < n; j+=1 {
		for i:=0; i < j; i+=1 {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j - i < 3 {   //不构成区间
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > maxlength {
				maxlength = j-i+1
				start, end = i, j
			}
		}
	}
	return s[start:end+1]
}


//方法3: 中心扩散
//枚举回文字符串中心, 求以每个回文中心得到的回文字符串的最大长度
func longestPalindromeExpandCenter(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	begin, end := 0, 0
	maxlength := 0
	var ExpandFromCenter func(string, int, int) (int, int)
	ExpandFromCenter = func(s string, left int, right int) (int, int) {  //[left, right]
		for {
			if left >=0 && right < n && s[left] == s[right]{
				left, right = left-1, right+1
			} else {
				break
			}
		}
		return right - left - 1, left+1     //length  right-1-(left+1)+1
	}
	for i:=0; i < n; i+=1 {
		len1, l1 := ExpandFromCenter(s, i, i)
		len2, l2 := ExpandFromCenter(s, i, i+1)
		len3, l3 := len1, l1
		if len2 > len1 {
			len3, l3 = len2, l2
		}
		if len3 > maxlength {
			maxlength, begin = len3, l3
			end = l3 + maxlength - 1 // [begin, end]
		}
	}
	return s[begin:end+1]
}

//方法4: Manacher 算法

//func main()  {
//	s := "cbbd"
//	fmt.Println(longestPalindromeExpandCenter(s))
//}