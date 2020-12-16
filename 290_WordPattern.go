package leetcode_go

import "strings"

func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")
	if len(strs) != len(pattern) {
		return false
	}
	pTos := make(map[byte]string)
	sTop := make(map[string]bool)
	for i := 0; i < len(strs); i++ {
		if pTos[pattern[i]] == "" {
			if !sTop[strs[i]] {
				pTos[pattern[i]] = strs[i]
				sTop[strs[i]] = true
			} else {
				return false
			}
		} else if pTos[pattern[i]] != strs[i] {
			return false
		}
	}
	return true
}
