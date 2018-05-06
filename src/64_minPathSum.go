package main

import (
	"fmt"
)

func getMin(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

func minPathSum(grid [][]int) int {
	var row int = len(grid)
	if row < 1 {
		return 0
	}

	var clo int = len(grid[0])
	var value []int = make([]int, row*clo)

	for i := 0; i < row; i++ {
		for j := 0; j < clo; j++ {
			if i == 0 && j == 0 {
				value[0] = grid[i][j]
			} else if i == 0 {
				value[clo*i+j] = grid[i][j] + value[clo*i+j-1]
			} else if j == 0 {
				value[clo*i+j] = grid[i][j] + value[clo*(i-1)+j]
			} else {
				value[clo*i+j] = grid[i][j] + getMin(value[clo*i+j-1], value[clo*(i-1)+j])
			}
		}
	}
	return value[row*clo-1]
}

func main() {
	arr1 := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(arr1)) // 7

	arr2 := [][]int{{1, 3, 1}}
	fmt.Println(minPathSum(arr2)) // 5

	arr3 := [][]int{{1}, {1}, {4}}
	fmt.Println(minPathSum(arr3)) //6
}
