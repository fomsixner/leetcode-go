func isToeplitzMatrix(matrix [][]int) bool {
	n, m := len(matrix), len(matrix[0])
	for i := 0; i < n; i++ {
		x, y := i, 0
		for x+1 < n && y+1 < m {
			if matrix[x][y] != matrix[x+1][y+1] {
				return false
			}
			x++
			y++
		}
	}
	for i := 0; i < m; i++ {
		x, y := 0, i
		for x+1 < n && y+1 < m {
			if matrix[x][y] != matrix[x+1][y+1] {
				return false
			}
			x++
			y++
		}
	}
	return true
}

// 根据定义，当且仅当矩阵中每个元素都与其左上角相邻的元素（如果存在）相等时，该矩阵为托普利茨矩阵。因此，我们遍历该矩阵，将每一个元素和它左上角的元素相比对即可。。
// func isToeplitzMatrix(matrix [][]int) bool {
//     n, m := len(matrix), len(matrix[0])
//     for i := 1; i < n; i++ {
//         for j := 1; j < m; j++ {
//             if matrix[i][j] != matrix[i-1][j-1] {
//                 return false
//             }
//         }
//     }
//     return true
// }
