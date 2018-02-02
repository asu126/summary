package main

import (
	"fmt"
)

// Time Limit Exceeded
// func climbStairs(n int) int {
// 	var sum int = 0
// 	if n <= 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	if n > 2 {
// 		sum = climbStairs(n-1) + climbStairs(n-2)
// 	}
// 	return sum
// }

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	dp := make([]int, n+1, n+1)
	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n-1]
}

func main() {
	fmt.Println(climbStairs(44))
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}
