package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if strings.Contains(a, b) {
		return b
	} else if strings.Contains(b, a) {
		return a
	}
	return "Nothing is the same"
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KIJANG", "KI"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
	fmt.Println(Compare("JOKO", "BEJO"))
}
