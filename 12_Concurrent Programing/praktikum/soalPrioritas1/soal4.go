// saya memodifikasi soal sehingga menghasilkan program yang...
// berjalan dengan multiple go routine yang dimasukan ke dalam looping
// karena jika soal dikerjakan dengan 1 go routine maka tidak mungkin terjadi race condition
// sehingga penggunaan mutex tidak diperlukan
// saya membuat program dimana mutex diperlukan untuk menghindari race condition

package main

import (
	"fmt"
	"sync"
)

type Number struct {
	sync.WaitGroup
	sync.Mutex
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func (num *Number) counter(n int, ch chan int) {
	num.Lock()

	temp := factorial(n)
	ch <- temp

	num.Unlock()
	num.Done()
}

func main() {
	n := []int{4, 5, 12, 8}
	ch := make(chan int, len(n))
	var num Number

	for _, i := range n {
		num.Add(1)
		go num.counter(i, ch)
	}
	num.Wait()
	close(ch)

	i := 0
	for result := range ch {
		fmt.Printf("Faktorial %d\t: %d\n", n[i], result)
		i++
	}
}
