package main

import (
	"fmt"
)

func findDiagonalOrder(matrix [][]int) []int {
	var M, N int = len(matrix), 0
	var result []int
	if M < 1 {
		return result
	}

	N = len(matrix[0])
	var lineSum int = 0
	var i, j int

	for count := 0; count < M*N; {

		i = lineSum
		j = 0
		for i >= 0 {
			// fmt.Println(i, j)
			if i < M && j < N {
				// fmt.Println("1", i, j)
				result = append(result, matrix[i][j])
				count++
			}
			i--
			j++
		}
		lineSum++

		i = 0
		j = lineSum
		for j >= 0 {
			// fmt.Println(i, j)
			if i < M && j < N {
				// fmt.Println("2", i, j)
				result = append(result, matrix[i][j])
				count++
			}
			i++
			j--
		}
		lineSum++

		// fmt.Println("...", count)
	}

	return result
}

func main() {
	var arr1 [][]int = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(findDiagonalOrder(arr1))

	var arr2 [][]int = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(findDiagonalOrder(arr2))

	var arr3 [][]int = [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(findDiagonalOrder(arr3))

	var arr4 [][]int = [][]int{}
	fmt.Println(findDiagonalOrder(arr4))

	var arr5 [][]int = [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	fmt.Println(findDiagonalOrder(arr5))
}
