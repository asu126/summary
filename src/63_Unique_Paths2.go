package main

import (
	"fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var m, n int = len(obstacleGrid), len(obstacleGrid[0])
	var dp [][]int = make([][]int, m)

	for i := 0; i < m; i++ {
		var nn []int = make([]int, n)
		dp[i] = nn
	}

	var tag1 bool = false
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 || tag1 {
			tag1 = true
			dp[0][i] = 0
		} else {
			dp[0][i] = 1
		}
	}

	tag1 = false
	for j := 0; j < m; j++ {
		if obstacleGrid[j][0] == 1 || tag1 {
			tag1 = true
			dp[j][0] = 0
		} else {
			dp[j][0] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// if obstacleGrid[i][j-1] == 0 && obstacleGrid[i-1][j] == 0 {
			// 	dp[i][j] = dp[i][j-1] + dp[i-1][j]
			// } else if obstacleGrid[i][j-1] == 1 {
			// 	dp[i][j] = dp[i-1][j]
			// } else {
			// 	dp[i][j] = dp[i][j-1]
			// }
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}

			// fmt.Println(i, j, dp[i][j])
		}
	}

	return dp[m-1][n-1]
}

func main() {
	x := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(x)) //2
	x1 := [][]int{
		{0},
		{1},
		{0},
	}
	fmt.Println(uniquePathsWithObstacles(x1)) // 0
	x2 := [][]int{
		{0},
	}
	fmt.Println(uniquePathsWithObstacles(x2)) // 1
	x3 := [][]int{
		{1, 0},
	}
	fmt.Println(uniquePathsWithObstacles(x3)) //0
}
