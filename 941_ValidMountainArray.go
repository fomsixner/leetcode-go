package leetcode_go

//判断是否为山脉数组

func validMountainArray(A []int) bool {
	n := len(A)
	if n < 3 {
		return false
	}
	i := 1
	for ; i<n-1; i+=1 {
		if A[i] <= A[i-1] {
			break
		}
	}
	if i == n || i == 1{
		return false
	}
	for ; i<n; i+=1 {
		if A[i] >= A[i-1] {
			return false
		}
	}
	return true
}
