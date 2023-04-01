# Section 12 - Concurrent Programing

## Goroutine
adalah mini thread atau lightweight
- sequential : 
  > menjalankan line code secara urut dari atas kebawah
  
  > sebelum menjalankan task baru, task lama harus sudah selesai
- paraller : 
  > banyak task dapat dijalankan bersamaan
- concurrent : 
  >banyak task dapat dijalankan **independent** (terpecah) tergantung bagian mana yang selesai terlebih dahulu

<br/>
<p align="center">
<img title="flowchart-symbol" alt="flowchart-symbol"  width= "600px" src="https://lh4.googleusercontent.com/ZWoxcnEe63j5EK46Q1b2tJco3FsL8uK-q-uDkAV2hTtl5n0COuoOHYsJATTnbh-Bh6s=w2400">
</p>


## Channel
> Channel digunakan untuk menghubungkan goroutine satu dengan goroutine lain. Mekanisme yang dilakukan adalah serah-terima data lewat channel tersebut. Dalam komunikasinya, sebuah channel difungsikan sebagai pengirim di sebuah goroutine, dan juga sebagai penerima di goroutine lainnya. Pengiriman dan penerimaan data pada channel bersifat blocking atau synchronous.

> Channel ibarat gudang yang menjadi tempat transit barang.
```golang
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	sayHello := func(str string, n time.Duration) {
		time.Sleep(n * time.Second)
		result := 0
		for i := range str {
			result += i
		}
		ch <- result
	}

	go sayHello("abcdefgh", 5)             //harusnya: 28
	go sayHello("abcd", 3)                 //harusnya: 6
	go sayHello("abcdefghkahsahsahsas", 2) //harusnya: 190

	fmt.Println(<-ch) //output: 190
	fmt.Println(<-ch) //output: 6
	fmt.Println(<-ch) //output: 28

	// URUTAN JADI TERBALIK
	// KARENA DURASI TIME SLEEP BERBEDA
	// TIME SLEEP TERCEPAT YANG AKAN DI PRINT TERLEBIH DAHULU
}
```

### Ada dua macam penerapan channel di golang :
1. Unbuffered Channel
   ```golang
   // Unbuffered channel (kanal tak berbuffer) adalah jenis channel yang menunggu hingga goroutine penerima siap untuk menerima data sebelum goroutine pengirim dapat mengirimkan data ke channel. 
    package main

    import (
        "fmt"
        "time"
    )
    ```
    ```golang
    // VERSI: 
    func main() {
        ch := make(chan string)

        sayHello := func(str string) {
            ch<-str
            fmt.Println("done ",str)
        }

        go sayHello("joko")
        go sayHello("bima")

        // fmt.Println(<-ch)
        // fmt.Println(<-ch)

        time.Sleep(3*time.Second)
    }
    // OUTPUT : *blank
    // karena unbuffered itu menuggu penerima siap, baru bisa menjalankan pengiriman
    // tapi disini penerima tidak di buat, sehingga pengiriman tidak dijalankan
   ```
2. Buffered Channel
   ```golang
   // buffered channel (kanal berbuffer) adalah jenis channel yang memungkinkan goroutine pengirim untuk mengirimkan data ke channel tanpa harus menunggu goroutine penerima siap untuk menerima data tersebut.
    // Jumlah data yang dapat disimpan di dalam buffered channel ditentukan saat deklarasi channel.
    // Ketika buffered channel penuh, goroutine pengirim akan diblokir sampai ada ruang di dalam channel untuk menampung data yang akan dikirimkan
    package main

    import (
        "fmt"
        "time"
    )
   ```
   ```golang
    // VERSI: buffered-1
    func main() {
        ch := make(chan string, 2)	// <--- disini kapasitas
                                    // kapasitas: jumlah data max mampu di tampung channel

        sayHello := func(str string) {
            ch<-str
            fmt.Println("done ",str)
        }

        go sayHello("joko")
        go sayHello("bima")

        // fmt.Println(<-ch)
        // fmt.Println(<-ch)

        time.Sleep(3*time.Second)
    }
    // OUTPUT : done bima
    //    		done joko
    // karena buffered tidak perlu menunggu penerima siap
    // jadi disini pengiriman tetap dijalankan
    // meskipun penerima belum siap
   ```
   ```golang
   // VERSI: buffered-2
    func main() {
        ch := make(chan string, 2)	// <--- disini kapasitas
                                    // kapasitas: jumlah data max mampu di tampung channel

        sayHello := func(str string) {
            ch<-str
            fmt.Println("done ",str)
        }

        // pengiriman
        go sayHello("joko")
        go sayHello("bima")
        go sayHello("eko")

        // penerimaan
        // fmt.Println(<-ch)
        // fmt.Println(<-ch)

        time.Sleep(3*time.Second)
    }
    // OUTPUT : done eko
    //    		done joko
    // hanya mengirimkan 2 data sesuai jumlah kapasitasnya
    // disini bima tidak beruntung sehingga tidak dikirimkan
    // jika ingin mengirimkan bima, maka harus menjalankan penerimaan
    // karena setiap penerimaan maka kapasitas channel akan berkurang 1
    // sehingga memungkinkan bima untuk dikirimkan ke channel
   ```

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
    //VERSI: 1
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
    ```golang
    //VERSI: 2
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

    func factorial(n int) int {
        if n == 0 {
            return 1
        }
        return n * factorial(n-1)
    }

    func counter(n int, ch chan int, mt *sync.Mutex, wg *sync.WaitGroup) {
        mt.Lock()

        temp := factorial(n)
        ch <- temp

        mt.Unlock()
        wg.Done()
    }

    func main() {
        n := []int{4, 5}
        ch := make(chan int, 2)
        var wg sync.WaitGroup
        var mt sync.Mutex

        for _, i := range n {
            wg.Add(1)
            go counter(i, ch, &mt, &wg)
        }
        wg.Wait()
        close(ch)

        for result := range ch {
            fmt.Println(result)
        }
    }
    //OUTPUT: 120
    //        24
    ```

<hr>

## Get API
```golang
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(responseData))
    // fmt.Println()
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
    // fmt.Println(responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	for _, pokemon := range responseObject.Pokemon {
		fmt.Println(pokemon.Species.Name)
	}
}

```

