package main

import (
	"fmt"
)

// Dynamic Programming
// Approach : Top-Down

// memoization
func fibo(n int, memo *map[int]int) int {
	if len(*memo) >= n {
		fmt.Println("pake memo")
		return (*memo)[n]
	}
	result := generateFibo(n, memo)
	return result
}

// generate fibonacci
func generateFibo(n int, memo *map[int]int) int {
	if n <= 1 {
		return n
	}
	// Cek apakah hasil perhitungan sudah ada di memoization
	if val, ok := (*memo)[n]; ok {
		return val
	}
	// Jika belum ada, hitung bilangan Fibonacci dengan rekursi
	(*memo)[n] = generateFibo(n-1, memo) + generateFibo(n-2, memo)
	return (*memo)[n]
}

func main() {
	memo := map[int]int{}
	fmt.Println(fibo(0, &memo))
	fmt.Println(memo)
	fmt.Println(fibo(1, &memo))
	fmt.Println(memo)
	fmt.Println(fibo(2, &memo))
	fmt.Println(memo)
	fmt.Println(fibo(3, &memo))
	fmt.Println(memo)
	fmt.Println(fibo(4, &memo))
	fmt.Println(memo)
	fmt.Println(fibo(5, &memo))
	fmt.Println(memo)
	fmt.Println(fibo(6, &memo))
	fmt.Println(memo)
	fmt.Println(fibo(3, &memo))
	fmt.Println(memo)
	fmt.Println(fibo(7, &memo))
	fmt.Println(memo)
	fmt.Println(fibo(8, &memo))
	fmt.Println(memo)
	fmt.Println(fibo(9, &memo))
	fmt.Println(memo)
	fmt.Println(fibo(10, &memo))
}
