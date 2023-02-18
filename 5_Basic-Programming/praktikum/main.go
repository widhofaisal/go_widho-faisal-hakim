package main

import (
	"fmt"
	"alterra/golang/5_Basic-Programming/praktikum/soalPrioritas1"
	"alterra/golang/5_Basic-Programming/praktikum/soalPrioritas2"
	"alterra/golang/5_Basic-Programming/praktikum/soalEksplorasi"
)

func main(){

	fmt.Println("\n------------------------------------------")
	// SOAL PRIORITAS 1
	fmt.Println("SOAL PRIORITAS 1")
	
	fmt.Println("\nNo.1")
	fmt.Println("Luas = ",soalprioritas1.LuasTrapesium())
	fmt.Println("\nNo.2")
	fmt.Println(soalprioritas1.BilanganGanjilGenap())
	fmt.Println("\nNo.3")
	fmt.Println(soalprioritas1.GradeNilai())
	fmt.Println("\nNo.4")
	soalprioritas1.FizzBuzz()
	
	fmt.Println("\n\n------------------------------------------")
	// SOAL PRIORITAS 2
	fmt.Println("SOAL PRIORITAS 2")

	fmt.Println("\nNo.1")
	soalprioritas2.SegitigaAsterik()
	fmt.Println("\nNo.2")
	soalprioritas2.FaktorBilangan()
	
	fmt.Println("\n------------------------------------------")
	// SOAL EKSPLORASI
	fmt.Println("\nSOAL EKSPLORASI")
	fmt.Println("\nNo.1")
	fmt.Println(soalEksplorasi.Palindrome())
}