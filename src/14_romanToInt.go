package main

import (
	"fmt"
)

/*
罗马数字字母
罗马数字只有7个字母组成，每个字母代码的字如下
字母	M	D	C	L	X	V	I
代表数字	1000	500	100	50	10	5	1

四个规则
相同的数字连写， 所表示的数等于这些数字相加得到的数。如 XXX表示 30
小的数字在大的数字的右边， 所表示的数等于这些数字相加得到的数 如VIII 表示8
小的数字(限于I, X, C)在大的数字的左边， 所表示的数等于大数减去小的数： 如IV 表示4
在一个数的上面画一条横线， 表示这个数增值1000倍(由于题目只考虑4000以内的数，所以这条规则不用考虑)。

*/

func romanToInt(s string) int {
	var tagVal [256]int
	tagVal['I'] = 1
	tagVal['V'] = 5
	tagVal['X'] = 10
	tagVal['C'] = 100
	tagVal['M'] = 1000
	tagVal['L'] = 50
	tagVal['D'] = 500
	var result int = 0

	for i := 0; i < len(s); i++ {
		if i+1 >= len(s) || tagVal[s[i+1]] <= tagVal[s[i]] {
			result += tagVal[s[i]]
		} else {
			result -= tagVal[s[i]]
		}

	}
	return result
}

func main() {
	fmt.Println(romanToInt("DCXXI"))
}
