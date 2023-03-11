package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10) //
	n := 3

	// send
	go func(n int) {
		for i := n; i < 100; i++ {
			if i%n == 0 {
				ch <- i
			}
		}
		close(ch)
	}(n)

	// receive
	go func() {
		for item := range ch {
			fmt.Println(item)
		}
	}()

	// waiting time
	time.Sleep(3 * time.Second)
}
