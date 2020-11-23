package leetcode_go

//字母只有26个,直接使用int数组就行了
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)
	for i, _ := range s {
		m[s[i]]++
	}
	for i, _ := range t {
		m[t[i]]--
		if m[t[i]] < 0 {
			return false
		}
	}
	return true
}
