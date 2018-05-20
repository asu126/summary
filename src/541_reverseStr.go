package main

import (
	"fmt"
)

func reverseStr(s string, k int) string {
	if k < 2 {
		return s
	}

	var length int = len(s)
	bs := []byte(s)

	var shang, mod = length / (2 * k), length % (2 * k)

	var count, i, j int = 0, 0, 0
	for ; count < shang; count++ {
		i = count * 2 * k
		j = i + k - 1
		for i < j {
			bs[i], bs[j] = bs[j], bs[i] // swap
			i++
			j--
		}
	}

	if mod > 0 {
		i = count * 2 * k
		if mod >= k {
			j = i + k - 1
		} else {
			j = i + mod - 1
		}
		for i < j {
			bs[i], bs[j] = bs[j], bs[i] // swap
			i++
			j--
		}
	}

	return string(bs)
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2))

	fmt.Println(reverseStr("abcdefgh", 2))

	fmt.Println(reverseStr("abcdefghi", 2))
}
