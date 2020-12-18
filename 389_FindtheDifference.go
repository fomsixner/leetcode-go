package leetcode_go

func findTheDifference(s string, t string) byte {
	m := make(map[byte]int)
	for i, _ := range s {
		m[s[i]]++
	}
	var res byte
	for i, _ := range t {
		m[t[i]]--
		if m[t[i]] < 0 {
			res = t[i]
			break
		}
	}
	return res
}
