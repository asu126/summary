package main

import (
	"fmt"
)

func judgeCircle(moves string) bool {
	var x, y int = 0, 0
	for i := 0; i < len(moves); i++ {
		// R（右），L（左），U（上）和 D（下
		switch moves[i] {
		case 'R', 'r':
			y++
		case 'L', 'l':
			y--
		case 'U', 'u':
			x--
		case 'D', 'd':
			x++
		}
	}

	if x == 0 && y == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(judgeCircle("UD"))
	fmt.Println(judgeCircle("LL"))
}
