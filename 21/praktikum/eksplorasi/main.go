package main

import "alterra/golang/21/praktikum/eksplorasi/route"

func main() {
	e := route.New()

	e.Start(":8000")
}