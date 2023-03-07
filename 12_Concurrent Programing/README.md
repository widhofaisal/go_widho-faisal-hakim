# Section 12 - Concurrent Programing

## Big Search Website
- sequential : sebelum menjalankan task baru, task lama harus sudah selesai
- paraller : banyak task dapat dijalankan bersamaan
- concurrent : banyak task dapat dijalankan **independent** (terpisah) dan memungkinkan secara **simultaneous** (bersamaan) 

## Goroutine

## Channel
- Example 1
    ```golang
    package main

    import (
        "fmt"
    )

    func greet(c chan string) {
        data := <-c
        fmt.Println("Hello " + data + " !")
    }

    func main() {
        fmt.Println("Main() started")
        c := make(chan string)

        go greet(c)

        c <- "john"
        fmt.Println("Main() stopped")
    }

    ```

- Example 2  
    ```golang
    package main

    import (
        "fmt"
        "time"
    )

    func main() {
        theMine := [5]string{"ore1", "ore2", "ore3"}
        oreChan := make(chan string)

        // Finder
        go func(mine [5]string) {
            for _, item := range mine {
                oreChan <- item
            }
        }(theMine)

        // Ore breaker
        go func() {
            for i := 0; i < 3; i++ {
                foundOre := <-oreChan
                fmt.Println("Miner: Received ", foundOre, " from finder")
            }
        }()

        <-time.After(time.Second * 5)
    }

    ```

### Ada dua macam penerapan channel di golang :
1. Buffered Channel
   ```golang
    package main

    import (
        "fmt"
    )

    func main() {
        message := make(chan string, 2)

        message <- "joko"
        message <- "wawan"

        fmt.Println(<-message)
        fmt.Println(<-message)
    }
   ```
2. Unbuffered Channel

<br/>

## Select
```golang
package main

import (
	"fmt"
	"time"
)

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, v := range numbers {
		sum += v
	}
	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, v := range numbers {
		if max < v {
			max = v
		}
	}
	// time.Sleep(1*time.Second)
	ch <- max
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// channel 1
	chan1 := make(chan float64)
	go getAverage(numbers, chan1)

	// channel 2
	chan2 := make(chan int)
	go getMax(numbers, chan2)

	select {
	case avg := <-chan1:
		fmt.Printf("Avg \t: %.2f \n", avg)
	case max := <-chan2:
		fmt.Printf("Max \t: %d \n", max)
	}
}

```

## Race Condition
How to check if there is race condition:
```golang
package main

import (
	"fmt"
)

func getNumber() int {
	i := 0
	go func() {
		i = 5
	}()
	return i
}

func main() {
	fmt.Println(getNumber())
}

// output: 0
// check of there is a race condition: go run -race main.go
```

How to solve race condition :
- Waitgroup
    ```golang
    package main

    import (
        "fmt"
        "sync"
    )

    func getNumber() int {
        i := 0
        var wg sync.WaitGroup

        wg.Add(1)
        go func() {
            i = 5
            wg.Done()
        }()
        wg.Wait()

        return i
    }

    func main() {
        fmt.Println(getNumber())
    }

    // output: 5
    ```

- Blocking
    ```golang
    package main

    import (
        "fmt"
        // "sync"
    )

    func getNumber() int {
        i := 0
        // var wg sync.WaitGroup
        ch := make(chan interface{})

        // wg.Add(1)
        go func() {
            i = 5
            // wg.Done()
            ch <- struct{}{}
        }()
        // wg.Wait()
        <-ch

        return i
    }

    func main() {
        fmt.Println(getNumber())
    }

    //output: 5
    ```

- Mutex
    ```golang
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

    //output: 5
    ```

