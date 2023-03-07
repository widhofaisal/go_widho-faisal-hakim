package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeNumber struct {
	i int
	m sync.Mutex
}

func (safeNumber *SafeNumber) Set(number int) {
	safeNumber.m.Lock()
	defer safeNumber.m.Unlock()
	safeNumber.i = number
}

func (safeNumber *SafeNumber) Get() int {
	safeNumber.m.Lock()
	defer safeNumber.m.Unlock()
	return safeNumber.i
}

func GetNumber() int {
	i := SafeNumber{}
	go func() {
		i.Set(5)
	}()
	time.Sleep(100 * time.Millisecond)
	return i.Get()
}

func main() {
	fmt.Println(GetNumber())
}
