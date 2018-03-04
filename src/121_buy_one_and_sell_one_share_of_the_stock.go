package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	var profit int = 0
	var length int = len(prices)
	var tmpMaxPrice int
	for i := 0; i < length; i++ {
		tmpMaxPrice = prices[i]
		for j := i + 1; j < length; j++ {
			if prices[j]-prices[i] > 0 {
				// fmt.Println("...", j)
				if prices[j] >= tmpMaxPrice {
					tmpMaxPrice = prices[j]
				}
			} else {
				break
			}
		}
		if tmpMaxPrice-prices[i] > profit {
			profit = tmpMaxPrice - prices[i]
		}
	}
	return profit
}

func main() {
	array := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(array))
}
