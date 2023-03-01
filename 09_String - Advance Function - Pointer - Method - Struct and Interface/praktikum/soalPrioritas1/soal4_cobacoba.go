package main

import (
	"fmt"
)

func getMinMax(min *int, max *int, numbers ...*int) {
	*min = *numbers[0]
	*max = *numbers[0]
	for _, item := range numbers {
		if *max < *item {
			*max = *item
		}
		if *min > *item {
			*min = *item
		}
	}
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
	getMinMax(&min, &max, &a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
}
