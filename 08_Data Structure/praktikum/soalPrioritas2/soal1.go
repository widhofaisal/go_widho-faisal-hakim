package main

import (
	"fmt"
)

func PairSum(arr []int, target int) []int {
	for index1, number1 := range arr {
		for index2, number2 := range arr {
			if number1+number2 == target {
				return []int{index1, index2}
			}
		}
	}

	return []int{}
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}
