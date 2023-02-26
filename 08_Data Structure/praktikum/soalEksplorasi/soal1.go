package main

import "fmt"

func SelisihDiagonal(length int, arr []int) {
	// check arr length is valid == length^2
	if length*length == len(arr) {
		var diagonal1, diagonal2 int
		var temp1, temp2 []int

		for i := 0; i < len(arr); i += (length + 1) {
			temp1 = append(temp1, arr[i])
			diagonal1 += arr[i]
		}
		// fmt.Println(diagonal1)

		for i := length - 1; i < len(arr)-1; i += (length - 1) {
			temp2 = append(temp2, arr[i])
			diagonal2 += arr[i]
		}
		// fmt.Println(diagonal2)

		result := diagonal1 - diagonal2
		if result < 0 {
			result *= -1
		}

		// count the selisih
		fmt.Println("------------------------------------------")
		for i := 0; i < len(arr); i++ {
			fmt.Print(arr[i], "\t")
			if (i-(length-1))%length == 0 {
				fmt.Println("\n")
			}
		}
		fmt.Println("Ukuran matrix \t: ", length, "x", length)
		fmt.Println("Diagonal kiri \t: ", temp1, " = ", diagonal1)
		fmt.Println("Diagonal kanan \t: ", temp2, " = ", diagonal2)
		fmt.Printf("NILAI MUTLAK \t:  |%d - %d| = %d\n\n\n", diagonal1, diagonal2, result)
	} else {
		fmt.Println("------------------------------------------")
		fmt.Println("Maaf panjang arrayMatrix tidak sesuai\n\n")
	}
}

func main() {
	// SelisihDiagonal( ukuranMatrix, arrayMatrix)
	// ukuranMatrix 	: harus sama misal 3x3 atau 5x5 dst
	// arrayMatrix 		: len(arrayMatrix) harus sejumlah ukuranMatrix^2

	SelisihDiagonal(3, []int{1, 2, 3, 4, 5, 6, 9, 8, 9})
	SelisihDiagonal(3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	SelisihDiagonal(4, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	SelisihDiagonal(5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25})
	SelisihDiagonal(5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25})
}
