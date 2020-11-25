package leetcode_go

func sortString(s string) string {
	res := make([]byte, 0, len(s))
	var count [26]int //s只包含小写英文字母
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	isNotComplete := true
	for isNotComplete {
		isNotComplete = false
		for i := 0; i < 26; i++ {
			if count[i] > 0 {
				count[i]--
				res = append(res, byte('a'+i))
				isNotComplete = true
			}
		}
		for i := 25; i >= 0; i-- {
			if count[i] > 0 {
				count[i]--
				res = append(res, byte('a'+i))
				isNotComplete = true
			}
		}
	}
	return string(res)
}
