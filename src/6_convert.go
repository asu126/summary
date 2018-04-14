package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	var length int = len(s)
	var result []byte = make([]byte, length)
	var mum int = numRows + (numRows - 2)
	if numRows <= 1 {
		return s
	}
	var i, j, k, current int = 0, 0, 0, 0
	for i = 0; i < numRows; i++ {
		k = mum - i
		if i == 0 || i == numRows-1 {
			for j = i; j < length; {
				result[current] = s[j]
				current++
				j = j + mum
			}
		} else {
			for j = i; j < length; {
				result[current] = s[j]
				current++
				if k < length {
					result[current] = s[k]
					k += mum
					// fmt.Println(byte(s[j+numRows-1]))
					current++
				}

				j = j + mum
			}
		}
	}
	return string(result)
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("ABCD", 2))
	fmt.Println(convert("ABCDEFGH", 4))
	fmt.Println(convert("A", 1))
	// fmt.Println(convert("abcabcbb"))
}
