package main

import (
	"fmt"
	"strconv"
)

func intToBinary(n int) []int {
	// inisialisasi tabel
	table := make([]int, n+1)

	for i := 0; i <= n; i++ {
		fmt.Println("i",i)
		
		j:=i
		var binary string
		for j > 0 {
			// Mendapatkan nilai modulus 2 dari bilangan bulat n
			digit := j % 2
			// Konversi digit ke string dan gabungkan ke variabel binary
			binary = fmt.Sprintf("%d%s", digit, binary)
			// Bagi n dengan 2 untuk mendapatkan bilangan bulat baru
			j /= 2
		}

		temp3,_:=strconv.Atoi(binary)
		table = append(table, temp3)
	}

	return table[:n+1]
}

func main() {
	fmt.Println(intToBinary(6))
	fmt.Println(intToBinary(3))
}
