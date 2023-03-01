package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min int, max int) {
	tempMin := *numbers[0]
	tempMax := *numbers[0]
	for _, item := range numbers {
		if tempMax < *item {
			tempMax = *item
		}
		if tempMin > *item {
			tempMin = *item
		}
	}
	return tempMin, tempMax
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
}
