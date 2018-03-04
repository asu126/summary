package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	var placeNum int = 0
	var length int = len(flowerbed)
	var before, after int
	for i := 0; i < length; i++ {
		if flowerbed[i] == 0 {
			if i == 0 {
				before = 0
			} else {
				before = flowerbed[i-1]
			}

			if i == length-1 {
				after = 0
			} else {
				after = flowerbed[i+1]
			}

			if before == 0 && after == 0 {
				flowerbed[i] = 1
				placeNum++
			}
		}
	}

	if placeNum >= n {
		return true
	} else {
		return false
	}
}

func main() {
	// array := []int{1, 0, 0, 0, 1}
	array := []int{0, 0, 1, 0, 0}
	fmt.Println(canPlaceFlowers(array, 1))
}
