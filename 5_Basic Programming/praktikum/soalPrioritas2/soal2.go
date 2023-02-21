package soalprioritas2

import "fmt"

func FaktorBilangan() {
	var N int
	fmt.Print("Masukan bilangan : ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		if N%i == 0 {
			fmt.Println(i)
		}
	}
}
