package main

import (
	"fmt"
)

func numJewelsInStones(J string, S string) int {
	var result int = 0

	for i := 0; i < len(S); i++ {
		for j := 0; j < len(J); j++ {
			if S[i] == J[j] {
				result++
				break
			}
		}
	}

	return result
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))
}
