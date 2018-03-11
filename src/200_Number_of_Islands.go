package main

import "fmt"

func numIslands(grid [][]byte) int {
	var max int = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				if dfs(grid, i, j) {
					max++
				}
			}
		}
	}
	return max
}

func dfs(grid [][]byte, i int, j int) bool {
	var m, n int = len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return true
	}

	if grid[i][j] == '0' {
		return true
	}
	grid[i][j] = '0'
	return dfs(grid, i+1, j) && dfs(grid, i-1, j) && dfs(grid, i, j+1) && dfs(grid, i, j-1)
}

func main() {
	/* 数组 - 5 行 2 列*/

	// var a [][]byte = [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	var a [][]byte = [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '1'}}

	fmt.Println(numIslands(a))
	/* 输出数组元素 */
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
