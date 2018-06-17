package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	var mapRansomNote map[byte]int = make(map[byte]int)
	var mapMagazine map[byte]int = make(map[byte]int)

	for i := 0; i < len(ransomNote); i++ {
		mapRansomNote[ransomNote[i]]++
	}

	for i := 0; i < len(magazine); i++ {
		mapMagazine[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		if mapMagazine[ransomNote[i]] < mapRansomNote[ransomNote[i]] {
			return false
		}
	}

	return true
}

func main() {
	// 	canConstruct("a", "b") -> false
	// canConstruct("aa", "ab") -> false
	// canConstruct("aa", "aab") -> true
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
