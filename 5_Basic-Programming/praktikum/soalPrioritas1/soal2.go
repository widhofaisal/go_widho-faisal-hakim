package soalprioritas1

import "fmt"

func BilanganGanjilGenap() string {
	var N int
	fmt.Print("Masukan bilangan : ")
	fmt.Scan(&N)

	if N%2 == 0 {
		return "Bilangan Genap"
	} else {
		return "Bilangan Ganjil"
	}
}
