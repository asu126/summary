package main

import "fmt"

func findComplement(num int) int {
	var temp int = 1 << 31

	var i int = 0
	for ; i < 32; i++ {
		if num&temp > 0 {
			break
		}
		temp = temp >> 1
	}
	// fmt.Println("--", i)
	for ; i < 32; i++ {
		// fmt.Println(temp)
		if num&temp > 0 {
			num = num & (^temp)
		} else {
			num = num | temp
		}
		temp = temp >> 1
	}

	return num
}

func main() {
	// fmt.Println(5 & (1 << 2))
	fmt.Println(findComplement(5)) // 2
	// fmt.Println(findComplement(1)) // 0
}
