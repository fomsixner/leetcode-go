package leetcode_go

func match(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var m [26]int
	for i := range s1 {
		m[s1[i]-'a']++
	}
	for i := range s2 {
		m[s2[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if m[i] != 0 {
			return false
		}
	}
	return true
}

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	isGroup := make([]bool, len(strs))
	for i := 0; i < len(strs); i++ {
		if !isGroup[i] {
			isGroup[i] = true
			var str []string
			str = append(str, strs[i])
			for j := i + 1; j < len(strs); j++ {
				if match(strs[i], strs[j]) {
					str = append(str, strs[j])
					isGroup[j] = true
				}
			}
			res = append(res, str)
		}
	}
	return res
}

//解法二
//func groupAnagrams(strs []string) [][]string {
//	mp := map[[26]int][]string{}
//	for _, str := range strs {
//		cnt := [26]int{}
//		for _, b := range str {
//			cnt[b-'a']++
//		}
//		mp[cnt] = append(mp[cnt], str)
//	}
//	ans := make([][]string, 0, len(mp))
//	for _, v := range mp {
//		ans = append(ans, v)
//	}
//	return ans
//}
