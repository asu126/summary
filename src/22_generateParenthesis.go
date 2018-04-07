package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	var result []string = make([]string, 0, 10240)
	getgenerateParent(n, 0, 0, "", &result)
	return result
}

func getgenerateParent(n int, leftRem int, rightRem int, s string, result *[]string) {
	// fmt.Println(n, leftRem, rightRem)
	if leftRem == n && rightRem == n {
		tmp := len(*result)
		*result = (*result)[:tmp+1]
		(*result)[tmp] = s
		// fmt.Println(s)
	} else {
		if leftRem < n { // 还有左括号可用则加入左括号
			// s += "("
			getgenerateParent(n, leftRem+1, rightRem, s+"(", result)
		}
		if leftRem > rightRem { // 左括号比右括号多就可以加入右括号
			// s += ")"
			getgenerateParent(n, leftRem, rightRem+1, s+")", result)
		}
	}
}
func main() {
	fmt.Println(generateParenthesis(0))
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
}
