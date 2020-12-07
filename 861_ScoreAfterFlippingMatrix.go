package leetcode_go

//首先将第一列都转换为1
//对剩下的列，仅做列转换，使每列的1尽可能多
//实际上并不需要翻转,只需记录每一列1的个数
func matrixScore(A [][]int) int {
	row, col := len(A), len(A[0])
	res := row * (1 << (col - 1))
	for j := 1; j < col; j++ {
		countOne := 0
		for i := 0; i < row; i++ {
			if A[i][0] == 1 {
				countOne += A[i][j]
			} else { //这一行经过了翻转
				countOne += 1 - A[i][j]
			}
		}
		if countOne < (row - countOne) {
			countOne = row - countOne
		}
		res += countOne << (col - j - 1)
	}
	return res
}
