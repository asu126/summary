package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	var lengthA, lengthB int = len(a), len(b)
	maxlen := lengthB
	if lengthA > lengthB {
		// minLen = lengthB
		maxlen = lengthA
	}

	var chars []byte = make([]byte, maxlen+1)
	var i, j, add int = lengthA - 1, lengthB - 1, 0

	for i >= 0 && j >= 0 {
		// fmt.Println(int(a[i]-'0'), int(b[j]-'0'), add)
		chars[maxlen] = byte((int(a[i]-'0')+int(b[j]-'0')+add)%2) + '0'
		add = (int(a[i]-'0') + int(b[j]-'0') + add) / 2
		i--
		j--
		maxlen--
	}

	for i >= 0 {
		chars[maxlen] = byte((int(a[i]-'0')+add)%2) + '0'
		add = (int(a[i]-'0') + add) / 2
		i--
		maxlen--
	}
	for j >= 0 {
		chars[maxlen] = byte((int(b[j]-'0')+add)%2) + '0'
		add = (int(b[j]-'0') + add) / 2
		j--
		maxlen--
	}

	if add > 0 {
		chars[0] = byte(add) + '0'
	} else {
		chars = chars[1:]
	}
	// fmt.Println(lengthA, lengthB, string(chars))
	return string(chars)
}

func main() {
	fmt.Println(addBinary("11", "10"))
	fmt.Println(addBinary("11", ""))
	fmt.Println(addBinary("11", "1"))
}
