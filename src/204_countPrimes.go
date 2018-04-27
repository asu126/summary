package main

import (
	"fmt"
)

func is_Primes(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	var notPrimes []bool = make([]bool, n)
	notPrimes[0] = true
	notPrimes[1] = true
	var count int = 0

	for i := 2; i < n; i++ {
		// if !notPrimes[i] && is_Primes(i) {
		if !notPrimes[i] {
			for j := i + i; j < n; j = j + i {
				notPrimes[j] = true
			}

		}
	}

	// fmt.Println(notPrimes)
	for i := 2; i < n; i++ {
		if !notPrimes[i] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(is_Primes(2))
	fmt.Println(countPrimes(1))
	fmt.Println(countPrimes(3))
	fmt.Println(countPrimes(5))
	fmt.Println(countPrimes(7))
	fmt.Println(countPrimes(8))
	fmt.Println(countPrimes(9))
	fmt.Println(countPrimes(11))
	fmt.Println(countPrimes(12))
	fmt.Println(countPrimes(499979))
	fmt.Println(countPrimes(12))
}
