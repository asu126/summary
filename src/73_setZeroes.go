package main

import (
	"fmt"
)

func setZeroes(matrix [][]int) {
	var lengthM int = len(matrix)
	if lengthM < 1 {
		return
	}

	var lengthN int = len(matrix[0])
	var zeorM, zeorN []bool = make([]bool, lengthM), make([]bool, lengthN)

	for i := 0; i < lengthM; i++ {
		for j := 0; j < lengthN; j++ {
			if matrix[i][j] == 0 {
				zeorM[i] = true
				zeorN[j] = true
			}
		}
	}

	for i := 0; i < lengthM; i++ {
		if zeorM[i] {
			for j := 0; j < lengthN; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for i := 0; i < lengthN; i++ {
		if zeorN[i] {
			for j := 0; j < lengthM; j++ {
				matrix[j][i] = 0
			}
		}
	}

}

func main() {
	// 	 [
	//   [1,1,1],
	//   [1,0,1],
	//   [1,1,1]
	// ]
	var arr1 [][]int = [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(arr1)
	fmt.Println(arr1)

	var arr2 [][]int = [][]int{{1, 1, 1}, {1, 0, 1}}
	setZeroes(arr2)
	fmt.Println(arr2)

	// [
	//   [0,1,2,0],
	//   [3,4,5,2],
	//   [1,3,1,5]
	// ]
	var arr3 [][]int = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(arr3)
	fmt.Println(arr3)
}
