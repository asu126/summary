package main

import (
	"fmt"
)

func convertToTitle(n int) string {
	bytes := make([]byte, 0, 100)

	for n > 0 {
		// 难道仅仅是因为从1开始计数的就要每次减一吗？？？
		// 1 -> A
		n -= 1
		bytes = append(bytes, byte(65+n%26))

		n /= 26
		// fmt.Println("...", n)
	}

	revBytes := make([]byte, len(bytes))
	for i, j := len(bytes)-1, 0; i >= 0; i-- {
		revBytes[i] = bytes[j]
		j++
	}

	return string(revBytes)
}

func main() {
	fmt.Println(int('A'))
	// fmt.Println(int('Z'))
	// fmt.Println(string(2 + 64))
	// fmt.Println(byte(26))

	// fmt.Println(convertToTitle(1))
	// fmt.Println(convertToTitle(28))
	// fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(520001))
}
