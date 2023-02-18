package soalprioritas2

import(
	"fmt"
)

func SegitigaAsterik() {
	var N int
	fmt.Print("\nMasukan bilangan : ")
	fmt.Scan(&N)

	for x := 1; x <= N; x++ {
		for y := N; y >= x; y-- {
			fmt.Print(" ")
		}
		for z := 1; z <= x; z++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
