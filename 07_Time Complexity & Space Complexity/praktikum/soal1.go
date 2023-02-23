package main

import (
	"fmt"
)

// Buatlah solusi dengan kompleksitas lebih cepat dari O(n) / O(n/2)
// O(sqrt(n))
func primeNumber(n int) bool {
	//    fmt.Println("number", n)
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		//   fmt.Printf("i=%v, i*i=%v \n", i, i*i)
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}
