package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	var dp [][]int = make([][]int, m)

	for i := 0; i < m; i++ {
		var nn []int = make([]int, n)
		dp[i] = nn
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
			// fmt.Println(i, j, dp[i][j])
		}
	}

	return dp[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 2))
	// m, n := 3, 2
	// var dp1 [m][n]int
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; i++ {
	// 		dp1[i][j] = i * j
	// 	}
	// }
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; i++ {
	// 		fmt.Print(dp1[i][j])
	// 	}
	// 	fmt.Println("")
	// }

}
