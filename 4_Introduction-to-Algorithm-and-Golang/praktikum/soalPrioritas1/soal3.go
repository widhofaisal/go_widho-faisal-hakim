package soalprioritas1

import "fmt"

func GradeNilai() string {
	var N int
	fmt.Print("Masukan nilai : ")
	fmt.Scan(&N)

	var grade string
	if N > 80 {
		grade = "A"
	} else if N > 65 {
		grade = "B"
	} else if N > 50 {
		grade = "C"
	} else if N > 35 {
		grade = "D"
	} else if N > 0 {
		grade = "E"
	}
	return grade
}
