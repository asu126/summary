package main

import "fmt"

func compress(chars []byte) int {
	var length int = len(chars)
	var strat, end, result int = 0, 0, 0
	var before byte
	var table []byte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	for strat < length {
		before = chars[strat]
		for end = strat + 1; end < length; end++ {
			if before != chars[end] {
				break
			}
		}
		count := end - strat
		chars[result] = before
		result++

		if count > 1 {
			var temp []int = make([]int, 0)

			for count != 0 {
				temp = append(temp, count%10)
				count = count / 10
			}

			for i := len(temp) - 1; i >= 0; i-- {
				chars[result] = table[temp[i]]
				result++
			}
		}
		strat = end
	}

	return result
}

func main() {
	// ["a","a","b","b","c","c","c"]
	var array0 []byte = []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(array0)) //6
	fmt.Println(array0)

	var array []byte = []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	fmt.Println(compress(array)) // 4
	fmt.Println(array)

}
