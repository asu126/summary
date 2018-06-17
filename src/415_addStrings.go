package main

import "fmt"

// import "strings"

func addStrings(num1 string, num2 string) string {
	if num1 == "" || num2 == "" {
		return num1 + num2
	}

	var i, j = len(num1) - 1, len(num2) - 1

	var result []int = make([]int, 0)
	var carry int = 0
	for i >= 0 && j >= 0 {
		sum := int(num1[i]-'0') + int(num2[j]-'0') + carry
		// fmt.Println(sum)
		result = append(result, sum%10)
		// fmt.Println(result)
		carry = sum / 10
		i--
		j--
	}

	for i >= 0 {
		sum := int(num1[i]-'0') + carry
		result = append(result, sum%10)
		carry = sum / 10
		i--
	}

	for j >= 0 {
		sum := int(num2[j]-'0') + carry
		result = append(result, sum%10)
		carry = sum / 10
		j--
	}

	if carry > 0 {
		result = append(result, carry)
	}
	// fmt.Println(result)
	var strArray []byte = make([]byte, len(result))
	i = 0
	for k := len(result) - 1; k >= 0; k-- {
		strArray[i] = byte(result[k] + '0')
		i++
	}
	return string(strArray[:])
}

func main() {
	fmt.Println(addStrings("1", "9"))
	fmt.Println(addStrings("111", "111"))
	fmt.Println(addStrings("999", "1001"))
}
