package main

import "alterra/golang/22_middleware/praktikum/route"

func main() {
	e := route.New()

	e.Start(":8000")
}