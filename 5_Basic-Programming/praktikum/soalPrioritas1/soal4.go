package soalprioritas1

import "fmt"

func FizzBuzz() {
	var output string
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			output = "FizzBuzz"
			fmt.Print(output, " ")
			continue
		} else if i%3 == 0 {
			output = "Fizz"
			fmt.Print(output, " ")
			continue
		} else if i%5 == 0 {
			output = "Buzz"
			fmt.Print(output, " ")
			continue
		}
		fmt.Print(i, " ")
	}
}
