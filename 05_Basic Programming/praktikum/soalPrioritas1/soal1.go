package soalprioritas1

import (
	"fmt"
)

func LuasTrapesium() float32 {
	var AB, CD, t float32

	fmt.Print("Sisi atas : ")
	fmt.Scan(&AB)

	fmt.Print("Sisi bawah : ")
	fmt.Scan(&CD)

	fmt.Print("Tinggi : ")
	fmt.Scan(&t)

	var Luas float32 = float32((AB + CD) / 2 * t)
	return Luas
}
