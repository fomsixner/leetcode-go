package leetcode_go

// 求岛屿周长
// 迭代法
// DFS 遍历法: 岛屿的周长就是岛屿方格和非岛屿方格相邻的边的数量。注意，这里的非岛屿方格，既包括水域方格，也包括网格的边界。
// 每当在 DFS 遍历中，从一个岛屿方格走向一个非岛屿方格，就将周长加 1

func islandPerimeter(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	res := 0
	for i, row := range grid {
		for j, v := range row {
			if v != 0 {
				if (i+1 < n && grid[i+1][j] == 0) || i+1 == n {
					res += 1
				}
				if (i-1 >= 0 && grid[i-1][j] == 0) || i == 0 {
					res += 1
				}
				if (j+1 < m && grid[i][j+1] == 0) || j+1 == m {
					res += 1
				}
				if (j-1 >= 0 && grid[i][j-1] == 0) || j == 0 {
					res += 1
				}
			}
		}
	}
	return res
}


func islandPerimeterUseDFS(grid [][]int) int {
	var dfs func(grid [][]int, i int, j int) int
	dfs = func(grid [][]int, i int, j int) int {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
			return 1               //碰到边界,将长度+1
		}
		if grid[i][j] == 0 {
			return 1               //碰到水域,将长度+1
		}
		if grid[i][j] != 1 {
			return 0               //已经遍历过
		}
		grid[i][j] = 2             //遍历
		return dfs(grid, i-1, j) + dfs(grid, i+1, j) + dfs(grid, i, j-1) + dfs(grid, i, j+1)
	}

	for i, row := range grid {
		for j, v := range row {
			if v != 0 {
				return dfs(grid, i, j)
			}
		}
	}
	//没有岛屿
	return 0
}