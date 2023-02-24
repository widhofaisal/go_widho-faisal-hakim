# Section 8 - Data Structure

## Video – Array (15:53)

Ciri-ciri **ARRAY** :
- hanya bisa menyimpan satu tipe data yang sama
- jumlah anggotanya fixed, tidak bisa di ubah

<br/>

## Video – Slice (14:12)

Ciri-ciri **SLICE** :
- size nya fleksibel
- tipe data ***reference*** artinya jika ada slice baru yang terbentuk dari slice lama, maka data elemen slice yang baru akan memiliki alamat memori yang sama dengan elemen slice lama. Setiap perubahan yang terjadi di elemen slice baru, akan berdampak juga pada elemen slice lama yang memiliki referensi yang sama.
    ```golang
    var slice1 =  []string{"apel", "bayam", "carrot"}
	var slice2 = slice1[:]

	fmt.Println(slice1)		// [apel bayam carrot]
	fmt.Println(slice2)		// [apel bayam carrot]

	slice1[1]="durian"

	fmt.Println(slice1)		// [apel durian carrot]
	fmt.Println(slice2)		// [apel durian carrot]
    ```
    ```golang
    fmt.Println(&slice1)		// &[apel bayam carrot]
	fmt.Println(&slice2)		// &[apel bayam carrot]
	
	fmt.Println(&slice1[0])		// 0xc0000241c0
	fmt.Println(&slice1[1])		// 0xc0000241c0
	fmt.Println(&slice1[2])		// 0xc0000241c0
    ```
  

Cara cek tipe data array atau slice :

```golang
var prime [5]int
var fruitsA = []string{"apple", "grape"}
mapHewan := map[string]int{}

fmt.Println(reflect.ValueOf(prime).Kind()) 		// array
fmt.Println(reflect.ValueOf(fruitsA).Kind())	// slice
fmt.Println(reflect.ValueOf(mapHewan).Kind())   // map
```

<br/>

## Video – Map (6:16)
> mirip dictionary di Python

<br/>

declarating map :
-   ```golang
    mapHewan := map[string]int{}
    ```
-   ```golang
	var mapHewan = map[string]int{}
    ```
-   ```golang
	var mapHewan map[string]int
    mapHewan = map[string]int{}
    ```
    agar tidak error perlu 2 line code sehingga lebih repot karena zero value dari map adalah nil, maka tiap variabel bertipe map harus di-inisialisasi secara explisit nilai awalnya (agar tidak nil).

<br/>

## Video – Function (18:39)
