package main

import (
	"fmt"
)

// func rotate(matrix [][]int) {
// 	var col int = len(matrix)
// 	if col < 1 {
// 		return
// 	}

// 	var row int = len(matrix[0])
// 	// 主对角线元素交换
// 	for i := 0; i < row; i++ {
// 		for x, y := 0, i; x < y; {
// 			tmp := matrix[x][y]
// 			matrix[x][y] = matrix[y][x]
// 			matrix[y][x] = tmp
// 			x++
// 			y--
// 		}
// 	}

// 	for i := 1; i < row; i++ {
// 		for x, y := row-1, i; x > y; {
// 			// fmt.Println(x, y)
// 			tmp := matrix[x][y]
// 			matrix[x][y] = matrix[y][x]
// 			matrix[y][x] = tmp
// 			x--
// 			y++
// 		}
// 	}

// 	// fmt.Println(matrix)
// 	// x元素交换

// 	for i := 0; i < row; i++ {
// 		for j := 0; j < row/2; j++ {
// 			tmp := matrix[j][i]
// 			matrix[j][i] = matrix[row-1-j][i]
// 			matrix[row-1-j][i] = tmp
// 		}
// 	}
// }

func rotate(matrix [][]int) {
	var col int = len(matrix)
	if col < 1 {
		return
	}

	var row int = len(matrix[0])
	// 主对角线元素交换
	for i := 0; i < row; i++ {
		for j := 0; j < row-i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[row-1-j][row-1-i]
			matrix[row-1-j][row-1-i] = tmp
		}
	}

	// fmt.Println(matrix)
	// x元素交换

	for i := 0; i < row; i++ {
		for j := 0; j < row/2; j++ {
			tmp := matrix[j][i]
			matrix[j][i] = matrix[row-1-j][i]
			matrix[row-1-j][i] = tmp
		}
	}
}

func main() {
	var a [][]int = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16}}
	rotate(a)
	fmt.Println(a)

	var a1 [][]int = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	rotate(a1)
	fmt.Println(a1)

	var a2 [][]int = [][]int{{}}
	rotate(a2)
	fmt.Println(a2)
}
