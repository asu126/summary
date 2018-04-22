package main

import (
	"fmt"
)

func combine(n int, k int) [][]int {
	var result [][]int
	var flag []bool = make([]bool, n)
	if n == k {
		var tmp []int = make([]int, k)
		for i := 0; i < n; i++ {
			tmp[i] = i + 1
		}
		result = append(result, tmp)
		return result
	}
	for i := 0; i < k; i++ {
		flag[i] = true
	}

	var startLocation int = k - 1
	var count int = 0

	for startLocation < n {
		var tmp []int = make([]int, k)
		count = 0
		for i := 0; i < len(flag); i++ {
			if flag[i] == true {
				tmp[count] = i + 1
				count++
				// fmt.Print(i + 1)
			}
		}
		result = append(result, tmp)
		// fmt.Println(tmp)
		// fmt.Println("--------")

		// change

		flag[startLocation] = false
		flag[startLocation+1] = true

		// 同时将其左边的所有“1”全部移动到数组的最左端
		move(flag, startLocation)

		startLocation = CheckoutOneZero(flag)
	}

	// print last
	count = 0
	var tmp []int = make([]int, k)
	for i := 0; i < len(flag); i++ {
		if flag[i] == true {
			tmp[count] = i + 1
			count++
			// fmt.Print(i + 1)
		}
	}
	// fmt.Println(tmp)
	result = append(result, tmp)

	return result
}

func move(b []bool, n int) {
	var count int = 0
	for i := 0; i < n; i++ {
		if b[i] {
			count++
		}
	}

	for i := 0; i < n; i++ {
		if i < count {
			b[i] = true
		} else {
			b[i] = false
		}
	}
}

func CheckoutOneZero(b []bool) int {
	for i := 1; i < len(b); i++ {
		if b[i-1] == true && b[i] == false {
			return i - 1
		}
	}
	return len(b)
}

func main() {
	fmt.Println(combine(5, 3))
	fmt.Println(combine(5, 1))
	fmt.Println(combine(1, 1))
	fmt.Println(combine(4, 2))
	fmt.Println(combine(2, 2))
}
