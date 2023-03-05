package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	result := []pair{}       // variabel for return
	temp := map[string]int{} // map for temporary variabel

	// store data to map
	// append the value when key matches
	for _, item := range items {
		if temp[item] == 0 {
			temp[item] = 1
		} else {
			temp[item] += 1
		}
	}

	// convert map to struct pair
	for key, value := range temp {
		result = append(result, pair{name: key, count: value})
	}

	// selection sort O(n^2)
	for range result {
		for i := 1; i < len(result); i++ {
			if result[i-1].count > result[i].count {
				result[i-1], result[i] = result[i], result[i-1]
			}
		}
	}

	return result
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang->1 ruby->2 js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// C->1 D->2 B->3 A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tennis"}))
	// football->1 basketball->1 tennis->1
}
