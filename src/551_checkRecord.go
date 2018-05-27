package main

import (
	"fmt"
)

func checkRecord(s string) bool {
	var length int = len(s)
	var countA int = 0
	var countL int = 0

	for i := 0; i < length; i++ {
		switch s[i] {
		case 'A':
			countA++
			countL = 0
		case 'L':
			countL++
		default:
			countL = 0
		}
		if countL > 2 {
			return false
		}
	}

	if countA < 2 {
		return true
	} else {
		return false
	}

}

func main() {
	fmt.Println(checkRecord("PPALLP"))  //true
	fmt.Println(checkRecord("PPALLL"))  //false
	fmt.Println(checkRecord("PPALLPL")) //true
	fmt.Println(checkRecord("PPALLAL")) //false
}
