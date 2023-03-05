package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	nameEncode string
	nameDecode string
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s student) keyGetter() string {
	// only modify the key
	//      abcdefghijklmnopqrstuvwxyz
	key := "kecbxjtworayfmqpusldnivhgz"
	return key
}

func (s *student) Encode() string {
	var nameEncode = ""

	// your code here
	key := s.keyGetter()
	target := strings.Split(s.name, "")

	for i := 0; i < len(s.name); i++ {
		index := strings.Index(key, target[i])
		if index >= 0 {
			move := rune(index) + 'a'
			nameEncode += string(move)

		} else {
			nameEncode += target[i]
		}
	}
	s.nameEncode=nameEncode
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

	// your code here
	alphabet :="abcdefghijklmnopqrstuvwxyz"
	key := s.keyGetter()
	target := strings.Split(s.name, "")

	for i := 0; i < len(s.name); i++ {
		index := strings.Index(alphabet, target[i])
		if index >= 0 {
			nameDecode += string(key[index])

		} else {
			nameDecode += target[i]
		}
	}

	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Println("\nEncode of student's name " + a.name + "is : " + c.Encode() + "\n")
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Println("\nDecode of student's name " + a.name + "is : " + c.Decode() + "\n")
	}
}
