package main

import (
	"fmt"
	"time"
)

func main() {
	x := 3

	go func(x int) {
		for i := x; i < 100; i++ {
			if i%x == 0 {
				fmt.Println(i)
			}
		}
	}(x)

	// waiting time
	time.Sleep(3 * time.Second)
}
