package main

import (
	"fmt"
)

/*
这种无限循环，最终都会出现4，故一旦遇见4即可知道进入了无限循环，判断是非快乐数字
*/

func isHappy(n int) bool {
	var result int = 0
	for {
		// fmt.Println(n)
		if n == 1 {
			return true
		}

		if n == 4 {
			return false
		}

		result = 0
		for n != 0 {
			result += (n % 10) * (n % 10)
			n = n / 10
			// fmt.Println(n, result)
		}
		n = result
	}
}

func main() {
	fmt.Println(isHappy(2))
	fmt.Println(isHappy(19))
}
