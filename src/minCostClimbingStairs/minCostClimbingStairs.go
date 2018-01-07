package main

import (
	"fmt"
)

func minCostClimbingStairs(cost []int) int {
	var length int = len(cost) + 1
	dp := make([]int, length, length)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i < length; i++ {
		dp[i] = getMinValue(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}
	return dp[length-1]
}

func getMinValue(x, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}

func main() {
	// 使用切片创建数组10, 15, 20
	// array := []int{0}
	// array := []int{10, 15, 20}
	// array := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	array := []int{1, 100}
	fmt.Println(array)
	fmt.Println(minCostClimbingStairs(array))
}
