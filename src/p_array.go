package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	// 注意*与谁结合，如p *[5]int，*与数组结合说明是数组的指针
	var p *[5]int = &a
	fmt.Println(*p, a)
	for index, value := range *p {
		fmt.Println(index, value)
	}
	for i := 0; i < len(*p); i++ {
		fmt.Println(i, (*p)[i])
	}

	// 如p [5]*int，*与int结合，说明这个数组都是int类型的指针，是指针数组。
	var p2 [5]*int
	i, j := 10, 20
	p2[0] = &i
	p2[1] = &j
	for index, value := range p2 {
		if value != nil {
			fmt.Println(index, *value)
		} else {
			fmt.Println(index, value)
		}
	}
	for i := 0; i < len(p2); i++ {
		if p2[i] != nil {
			fmt.Println(i, (p2[i]))
			fmt.Println(i, *(p2[i]))
		}
	}
}
