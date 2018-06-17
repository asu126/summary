package main

import "fmt"

func getSum(a int, b int) int {
	if a&b == 0 {
		return a | b
	}
	return getSum(a^b, (a&b)<<1)
}

func main() {
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(100, 200))
	fmt.Println(getSum(0, 2))
	fmt.Println(getSum(-1, 1))
	fmt.Println(getSum(-1, -1))
}
