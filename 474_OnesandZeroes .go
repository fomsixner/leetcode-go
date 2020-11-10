package leetcode_go

//1和0
//动态规划, 0-1背包问题


func findMaxForm(strs []string, m int, n int) int {
	length := len(strs)
	zeros := make([]int, length)
	ones := make([]int, length)
	for i, str:= range strs {    //获得每个字符串的0,1个数
		for _, c:= range str {
			if c == '1' {
				ones[i] += 1
			} else {
				zeros[i] += 1
			}
		}
	}
	dp := make([][]int, m+1)
	for i, _:= range dp {
		dp[i] = make([]int, n+1)
	}
	//初始化
	//dp[0][0] = 0   //可以不写, 因为默认赋值为0
	//for i:=0; i<=m; i+=1 {
	//	for j:=0; j<=n; j+=1 {
	//		for k:=0; k < length; k+=1 {
	//			if zeros[k] <= i && ones[k] <= j && (dp[i-zeros[k]][j-ones[k]]+1) > dp[i][j] {
	//				dp[i][j] = dp[i-zeros[k]][j-ones[k]] + 1
	//			}
	//		}
	//	}
	//}
	for k:=0; k<length; k+=1 {
		for i:=m; i>= zeros[k]; i-=1 {
			for j:=n; j>=ones[k]; j-=1 {
				if (dp[i-zeros[k]][j-ones[k]]+1) > dp[i][j] {
					dp[i][j] = dp[i-zeros[k]][j-ones[k]] + 1
				}
			}
		}
	}
	//for _,c:=range dp{
	//	for _,d:=range c{
	//		fmt.Printf("%d ", d)
	//	}
	//	fmt.Printf("\n")
	//}
	return dp[m][n]
}

//func main()  {
//	strs := []string{"10","0001","111001","1","0"}
//	fmt.Println(findMaxForm(strs, 5, 3))
//}