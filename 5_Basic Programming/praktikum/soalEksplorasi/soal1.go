package soalEksplorasi

import "fmt"

func Palindrome() string {
	line := "aya aya"
	fmt.Println("Masukan kata : ", line)

	for i := 0; i < len(line); i++ {
		j := len(line) - 1 - i
		if line[i] != line[j] {
			return "Bukan palindrome"
		}
	}

	return "Palindrome"
}
