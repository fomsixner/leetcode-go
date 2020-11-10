package leetcode_go

//湍流子数组

//超时
//func maxTurbulenceSize(A []int) int {
//	n, res:= len(A), 1
//	if n < 2 {    //只有1个数时一定为湍流数组
//		return n
//	}
//	//if n == 2 && A[]
//	dp := make([][]bool, n)
//	for i, _:=range dp {  //自动初始化为false
//		dp[i] = make([]bool, n)
//	}
//	//初始化,单个数一定为湍流数组
//	for i:=0; i<n; i++ {
//		dp[i][i] = true
//		if i+1 < n {
//			if A[i] != A[i+1]{
//				dp[i][i+1] = true
//				res = 2
//			} else {
//				dp[i][i+1] = false
//			}
//		}
//	}
//	for j:=2; j<n; j++ {
//		for i:=0; i<j-1; i++ {
//			dp[i][j] = dp[i][j-1] && ((A[j] > A[j-1] && A[j-1] < A[j-2]) || (A[j] < A[j-1] && A[j-1] > A[j-2]))
//			if dp[i][j] && j-i+1 > res {
//				res = j-i+1
//			}
//		}
//	}
//	return res
//}

//滑动窗口
func maxTurbulenceSize(A []int) int {
	var cmp func(int, int) int
	cmp = func(a int, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}
	n := len(A)
	left, res := 0, 1
	for j:=1; j<n; j++ {
		if j==n-1 || cmp(A[j-1],A[j])*cmp(A[j], A[j+1])!=-1 {
			if cmp(A[j-1],A[j])!=0 && j-left+1 > res {
				res = j-left+1
			}
			left = j
		}
	}
	return res
}


//func main()  {
//	nums := []int{0,1,1,0,1,0,1,1,0,0}
//	fmt.Println(maxTurbulenceSize(nums))
//}