package main

import "fmt"

func reverseSlice(s []int) {
    for i := 0; i < len(s)/2; i++ {
        j := len(s) - i - 1
        s[i], s[j] = s[j], s[i]
    }
    s = append(s, 6) // menambahkan elemen baru ke slice
}

func main() {
	arr:=[]int{1,2,3}
	fmt.Println(arr)
	reverseSlice(arr)
	fmt.Println(arr)
}

// urutan slice diluar fungsi dapat dipengaruhi perubahan slide didalam fungsi
// namun value nya tidak bisa dipengaruhi