package leetcode_go

// * / + -
func clumsy(N int) int {
	ans := 0
	var seq []int
	for i := 0; i < N; i++ {
		seq = append(seq, N-i)
	}
	for i := 0; i*4 < len(seq); i++ {
		if i*4+2 < len(seq) {
			if i == 0 {
				ans += seq[i*4] * seq[i*4+1] / seq[i*4+2]
			} else {
				ans -= seq[i*4] * seq[i*4+1] / seq[i*4+2]
			}
		} else if i*4+1 < len(seq) {
			if i == 0 {
				ans += seq[i*4] * seq[i*4+1]
			} else {
				ans -= seq[i*4] * seq[i*4+1]
			}
		} else {
			if i == 0 {
				ans += seq[i*4]
			} else {
				ans -= seq[i*4]
			}
		}
		if i*4+3 < len(seq) {
			ans += seq[i*4+3]
		}
	}
	return ans
}
