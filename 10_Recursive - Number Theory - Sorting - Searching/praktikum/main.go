package main

import "fmt"


func main() {
	for i:=1; i<100;i++{
		temp := []int{}
		for j:=2; j<=i; ++{
			if j%i==0{
				temp = append(temp, i)
			}
		}
		fmt.Println(temp)
	}
}
