package main

import (
	"fmt"
	"strconv"
	"strings"
)

func munculSekali(angka string) []int {
	arrayA := strings.Split(angka, "")
	mapA := map[string]int{}
	result := []int{}

	for _, item := range arrayA {
		mapA[item] += 1
		if mapA[item] > 1 {
			delete(mapA, item)
		}
	}

	for key, _ := range mapA {
		intVar, _ := strconv.Atoi(key)
		result = append(result, intVar)
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
