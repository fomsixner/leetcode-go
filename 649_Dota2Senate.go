package leetcode_go

//模拟循环队列
func predictPartyVictory(senate string) string {
	var radiant, dire []int //存储每个议员投票的次序
	for i, v := range senate {
		if v == 'D' {
			dire = append(dire, i)
		} else {
			radiant = append(radiant, i)
		}
	}
	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] {
			dire = dire[1:]
			radiant = append(radiant[1:], radiant[0]+len(senate))
		} else {
			radiant = radiant[1:]
			dire = append(dire[1:], dire[0]+len(senate))
		}
	}
	if len(radiant) == 0 {
		return "Dire"
	}
	return "Radiant"
}
