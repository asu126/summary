package main

import "fmt"
import "strconv"

func summaryRanges(nums []int) []string {
	var ret []string = make([]string, 0, len(nums))

	var i, j int
	for i = 0; i < len(nums); {
		for j = i + 1; j < len(nums); j++ {
			if nums[j]-nums[j-1] != 1 {
				break
			}
		}
		var str string
		if j-1 == i {
			str = strconv.Itoa(nums[i])
		} else {
			str = strconv.Itoa(nums[i]) + "->" + strconv.Itoa(nums[j-1])
		}
		tmp := len(ret)
		ret = ret[:tmp+1]
		// fmt.Println(str)
		ret[tmp] = str
		i = j
	}

	return ret
}

func main() {
	a := []int{1, 2, 3, 5, 7, 8}

	fmt.Println(summaryRanges(a)) // true
}
