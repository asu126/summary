package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	var length int = len(strs)
	if length < 1 {
		return ""
	}
	var i, j int
	for i = 0; i < len(strs[0]); i++ {
		for j = 1; j < length; j++ {
			if len(strs[j]) <= i {
				return strs[0][:i]
			}
			if strs[0][i] != strs[j][i] {
				if i < 1 {
					return ""
				} else {
					return strs[0][:i]
				}
			}
		}
	}

	return strs[0]
}

func main() {
	var ss1 []string = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(ss1))

	var ss2 []string = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(ss2))

	var ss3 []string = []string{"aa", "", "a"}
	fmt.Println(longestCommonPrefix(ss3))
}
