package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	result := []string{}
	arrayA = append(arrayA, arrayB...)
	mapA := map[string]int{}

	// convert array to map and elimination deplicates
	for _, item := range arrayA {
		mapA[item] = 1
	}

	// convert map to array
	for letter, _ := range mapA {
		result = append(result, letter)
	}

	return result
}

func main() {
	// Test cases

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese", "akuma"}))
	// ["king","devil jin","akuma","eddie","steve","geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei","jin","steve","bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa","yoshimitsu","devil jin","law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin","sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}
