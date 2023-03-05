package main

import "fmt"

// Dynamic Programming
// Approach : Top-Down

// generate binary according to n
func getBinary(n int, memo *[]string) []string {
	result := fmt.Sprintf("%b", n)
	*memo = append(*memo, result)
	if n == 0 {
		return *memo
	} else {
		getBinary(n-1, memo)
		return []string{}
	}
}

// using memoization
func setResult(n int, memo *[]string) []string {
	if len(*memo) >= n {
		return (*memo)[:n+1]
	} else {
		*memo = []string{}
		getBinary(n, memo)
		result := *memo
		return reverseSlice(result)
	}
}

// reverse the slice order
func reverseSlice(s []string) []string {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	memo := []string{}
	fmt.Println(setResult(2, &memo))
	fmt.Println(setResult(3, &memo))
	fmt.Println(setResult(7, &memo))
	fmt.Println(setResult(5, &memo))
	fmt.Println(setResult(4, &memo))
}
