package main

import (
	"fmt"
)

func minCostClimbingStairs(cost []int) int {
	var length, i int = len(cost), 0
	var sum int = 0

	if length == 0 {
		return 0
	}

	for i = 0; i < length; {
		// fmt.Println(i)
		switch length - i {
		case 1:
			sum += 0
			return sum
		case 2:
			if cost[i] > cost[i+1] {
				sum += cost[i+1]
			} else {
				sum += cost[i]
			}
			return sum
			// golang 默认添加了break
		case 3:
			if cost[i]+cost[i+2] > cost[i+1] {
				sum += cost[i+1]
			} else {
				sum += (cost[i] + cost[i+2])
			}
			return sum
		default:
			if cost[i+1] == 0 || (getMinValue(cost[i]+cost[i+1], cost[i]+cost[i+2]) >= getMinValue(cost[i+1]+cost[i+2], cost[i+1]+cost[i+3])) {
				sum += cost[i+1]
				i += 2
			} else {
				sum += cost[i]
				i += 1
			}
		}
	}

	return sum
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
	// array := []int{0, 1, 2, 2}
	array := []int{1, 0, 2, 2}
	fmt.Println(array)
	fmt.Println(minCostClimbingStairs(array))
}
