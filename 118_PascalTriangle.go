package leetcode_go

func generate(numRows int) [][]int {
	//var res [][]int
	//if numRows == 0 {
	//	return res
	//}
	//for i := 1; i <= numRows; i++ {
	//	if i == 1 {
	//		res = append(res, []int{1})
	//		continue
	//	}
	//	var row []int
	//	for j := 0; j < i; j++ {
	//		if j == 0 || j == i-1 {
	//			row = append(row, 1)
	//		} else {
	//			row = append(row, res[len(res)-1][j-1] + res[len(res)-1][j])
	//		}
	//	}
	//	res = append(res, row)
	//}
	//return res
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}

//func main()  {
//	generate(5)
//}
