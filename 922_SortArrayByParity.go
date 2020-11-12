package leetcode_go

//双指针一遍扫描
func sortArrayByParityII(A []int) []int {
	n := len(A)
	i, j := 0, 1
	for i < n && j < n {
		for ; i < n && A[i]%2 == 0; i += 2 {
		}
		for ; j < n && A[j]%2 == 1; j += 2 {
		}
		if i < n && j < n {
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}
