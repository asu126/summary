package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	var result []int
	if len(matrix) < 1 {
		return result
	}
	var m, n int = len(matrix), len(matrix[0])

	result = make([]int, m*n)

	var left, right, low, high int = 0, n - 1, 0, m - 1
	var number int = 0

	for left <= right && low <= high {
		for i := left; i <= right; i++ {
			// fmt.Println("1", left, right, low, high)
			// fmt.Println(matrix[low][i])
			result[number] = matrix[low][i]
			number++
		}
		low++

		for i := low; i <= high; i++ {
			// fmt.Println("2", left, right, low, high)
			// fmt.Println(matrix[i][right])
			result[number] = matrix[i][right]
			number++
		}
		right--

		for i := right; number < m*n && i >= left; i-- {
			// fmt.Println("3", left, right, low, high)
			// fmt.Println(matrix[i][left])
			result[number] = matrix[high][i]
			number++
		}
		high--

		for i := high; number < m*n && i >= low; i-- {
			// fmt.Println("4", left, right, low, high)
			// fmt.Println(matrix[i][left])
			result[number] = matrix[i][left]
			number++
		}
		left++

	}
	// fmt.Println(result)
	return result
}

func main() {

	var a1 [][]int = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16}}
	fmt.Println(spiralOrder(a1))

	var a2 [][]int = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	fmt.Println(spiralOrder(a2))

	var a3 [][]int = [][]int{
		{1, 2},
		{4, 5},
		{7, 8}}
	fmt.Println(spiralOrder(a3))

	var a4 [][]int = [][]int{}
	fmt.Println(spiralOrder(a4))

	var a5 [][]int = [][]int{{2, 3}}
	fmt.Println(spiralOrder(a5))
	var a6 [][]int = [][]int{{2}, {3}, {4}}
	fmt.Println(spiralOrder(a6))

	var a7 [][]int = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12}}
	fmt.Println(spiralOrder(a7))
}
