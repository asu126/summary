package main

import "fmt"
import "strings"

func countSegments(s string) int {
	temp := strings.Split(s, " ")
	var count int = 0
	for i := 0; i < len(temp); i++ {
		if temp[i] != "" {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countSegments("  Hello, hello"))
}
