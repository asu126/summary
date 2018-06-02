package main

import (
	"fmt"
)

// int,err:=strconv.Atoi(string)
// #string到int64
// int64, err := strconv.ParseInt(string, 10, 64)
// #int到string
// string:=strconv.Itoa(int)

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int = make([]int, 0)
	var result int = 0

	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			result = stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		case "-":
			result = stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		case "*":
			result = stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		case "/":
			result = stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		default:
			temp, _ := strconv.Atoi(tokens[i])
			stack = append(stack, temp)
		}
	}

	if len(stack) > 0 {
		result = stack[0]
	}
	return result
}

func main() {
	var arr1 []string = []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(arr1)) // 9

	var arr2 []string = []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(arr2)) // 6

	var arr3 []string = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(arr3)) // 22

	var arr4 []string = []string{"18"}
	fmt.Println(evalRPN(arr4)) // 18
}
