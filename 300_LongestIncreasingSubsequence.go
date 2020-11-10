package leetcode_go

//最长上升子序列

//动态规划

//dp[i] 表示以nums[i]结尾的最大子串长度

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n==0 {
		return 0
	}
	dp := make([]int, n)
	for i:=0; i<n; i+=1 {
		dp[i] = 1             //初始化
		for j:=0; j<i; j+=1 {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j]+1
				}
			}
		}
	}
	res := dp[0]
	for _, c:= range dp {
		if c > res {
			res = c
		}
	}
	return res
}
//
//func main()  {
//	nums := []int{3,5,6,2,5,4,19,5,6,7,12}
//	fmt.Println(lengthOfLIS(nums))
//}