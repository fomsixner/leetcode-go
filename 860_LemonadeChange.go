package leetcode_go

func lemonadeChange(bills []int) bool {
	m := make(map[int]int)
	for _, bill := range bills {
		if bill == 5 {
			m[bill]++
			continue
		} else if bill == 10 {
			if m[5] == 0 {
				return false
			} else {
				m[5]--
				m[10]++
			}
		} else { //bill == 20
			temp := 15
			if m[10] > 0 {
				m[10]--
				temp -= 10
			}
			if m[5] >= temp/5 {
				m[5] -= temp / 5
			} else {
				return false
			}
		}
	}
	return true
}
