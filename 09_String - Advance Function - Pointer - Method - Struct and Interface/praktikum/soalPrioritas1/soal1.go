package main

import "fmt"

type Car struct {
	typeCar string
	fuelIn  float32
}

func (c Car) mileage() float32 {
	return c.fuelIn / 1.5
}

func main() {
	mobil1 := Car{"avanza", 56.5}
	fmt.Println("Bisa menempuh jarak sejauh:", mobil1.mileage(), "mill")

	mobil2 := Car{"agya", 120.4}
	fmt.Println("Bisa menempuh jarak sejauh:", mobil2.mileage(), "mill")
}
