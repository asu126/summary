package main

import (
	"fmt"
)

func mySqrt(x int) int {
	var l, h int = 0, x
	var mid int
	if x == 1 {
		return 1
	}
	if x < 1 {
		return 0
	}
	for l <= h {
		mid = l + (h-l)/2
		// fmt.Println("xxx", mid)
		if mid == x/mid {
			return mid
		} else if mid > x/mid {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return h
}

func main() {
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(2))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}
