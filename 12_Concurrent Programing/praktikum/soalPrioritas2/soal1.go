package main

import (
	"fmt"
	"strings"
	"time"
)

func letterFrequency(text string, ch chan map[rune]int) {
	count := map[rune]int{}
	for _, char := range text {
		count[char]++
	}

	ch <- count
}

func main() {
	ch := make(chan map[rune]int)
	text := "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Sed eos sunt iste recusandae, et corporis cumque quia impedit nostrum totam."
	
	go letterFrequency(strings.Join(strings.Split(text, " "), ""), ch)
	
	time.Sleep(1 * time.Second)

	result := <-ch
	for key, value := range result {
		// fmt.Println(key, value)
		fmt.Printf("%c = %d\n", key, value)
	}
}
