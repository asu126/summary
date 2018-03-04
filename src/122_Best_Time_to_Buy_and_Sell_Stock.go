package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	var profit int = 0
	var length int = len(prices)
	for i := 1; i < length; i++ {
		if prices[i]-prices[i-1] > 0 {
			profit += (prices[i] - prices[i-1])
		}
	}
	return profit
}

func main() {
	array := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(array))
}
