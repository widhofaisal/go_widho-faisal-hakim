package main

import "fmt"

func primeX(number int) int {
	result := []int{} // variable for return

	// first loop: do infinite loops
	for i := 2; i>0; i++ {
		temp := []int{}

		// second loop: do loops until i
		for j := 2; j <= i; j++ {
			if i%j == 0 {
				temp = append(temp, j)
			}
		}

		if len(temp) == 1 {
			result = append(result, i)
		}
		
		// do the brake when the length of the result is equal to the number
		if len(result) == number {
			break
		}
	}

	return result[number-1]
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
