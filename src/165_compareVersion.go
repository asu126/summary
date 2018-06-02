package main

import (
	"fmt"
)

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	var v1 []int = make([]int, 0)
	var v2 []int = make([]int, 0)

	temp1 := strings.Split(version1, ".")
	for i := 0; i < len(temp1); i++ {
		t, _ := strconv.Atoi(temp1[i])
		v1 = append(v1, t)
	}

	for i := (len(v1) - 1); i >= 0; i-- {
		if v1[i] == 0 {
			v1 = v1[:i]
		} else {
			break
		}
	}

	temp2 := strings.Split(version2, ".")
	for i := 0; i < len(temp2); i++ {
		t, _ := strconv.Atoi(temp2[i])
		v2 = append(v2, t)
	}

	for i := (len(v2) - 1); i >= 0; i-- {
		if v2[i] == 0 {
			v2 = v2[:i]
		} else {
			break
		}
	}

	var i int
	for i = 0; i < len(v1) && i < len(v2); i++ {
		if v1[i] > v2[i] {
			return 1
		} else if v1[i] < v2[i] {
			return -1
		}
	}

	if i < len(v1) {
		return 1
	}
	if i < len(v2) {
		return -1
	}

	return 0

}

func main() {
	fmt.Println(compareVersion("7.5.2.4", "7.5.3")) //-1

	// version1 = "0.1", version2 = "1.1"
	fmt.Println(compareVersion("0.1", "1.1")) //-1

	//  version1 = "1.0.1", version2 = "1"
	fmt.Println(compareVersion("1.0.1", "1")) //1
}
