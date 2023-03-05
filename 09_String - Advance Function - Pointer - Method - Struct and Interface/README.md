# Section 9 - String - Advance Function - Pointer - Method - Struct and Interface

## Video – String (20:17)

```golang
// comparison
fmt.Println("abc"=="Abc")	// false
fmt.Println("abc"=="abc")	// true
```

```golang
// contains
letter1 := "indonesia"
fmt.Println(strings.Contains(letter1, "indo")) 	// true
fmt.Println(strings.Contains(letter1, "ido"))	// false
fmt.Println(strings.Contains(letter1, "indO"))	// false
```

```golang
// slicing string
var letter1 string = "indonesia raya"
slice1:=letter1[:4] ; fmt.Println(slice1) 	// indo
slice2:=letter1[10:] ; fmt.Println(slice2) 	// raya
```

```golang
// replace string
letter1 := "iwak iwak iwak"
fmt.Println(strings.Replace(letter1, "wak", "kan", 1))      // ikan iwak iwak
fmt.Println(strings.Replace(letter1, "wak", "kan", 2))      // ikan ikan iwak
fmt.Println(strings.Replace(letter1, "wak", "kan", 4))      // ikan ikan ikan
fmt.Println(strings.Replace(letter1, "wak", "kan", -1))     // ikan ikan ikan
```

```golang
// INSERT STRING

// insert front
letter := "green"
item1 := "high"
letter = item1 + " "+ letter
fmt.Println(letter)     // high green

// insert beetwen
letter := "green mountain"
item1 := "high"
letter = letter[:6] + item1 + letter[5:]
fmt.Println(letter)		// green high mountain
```

```golang
// shuffle string
import (
	"fmt"
	"math/rand"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	runes := []rune(alphabet)

	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})

	result := string(runes)
	fmt.Println(result)
}
```

<br/>

## Video – Advance Function (23:53)
- Variadic Func
  
- Anonymous Func
    ```golang
    // anonymous func
	bagianAwal := func(nama string){
		fmt.Println("Bagian awal")
		fmt.Println("Selamat datang", nama)
	}

	fmt.Println("Bagian akhir")
	fmt.Println("Selamat jalan")

	bagianAwal("joko")
	bagianAwal("bejo")
    ```

- Closure
    > aku jujur masih bingung closure yang mana dan fungsinya apa dan urgentsi untuk dipake kenapa
    ```golang
    // VERSI-1
    func main() {
        var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

        var newNumbers = func(min int) []int {
            var r []int
            for _, e := range numbers {
                if e < min {
                    continue
                }
                r = append(r, e)
            }
            return r
        }(3)

        fmt.Println("original number :", numbers)
        fmt.Println("filtered number :", newNumbers)
    }
    ```
    ```golang
    // VERSI-2
    func NewCounter() func() int {
        count := 0
        return func() int {
            count += 1
            return count
        }
    }

    func main() {
        counter := NewCounter()
        fmt.Println(counter())
        fmt.Println(counter())
        fmt.Println(counter())
    }
    ```

- Defer func
    ```golang
    func openDB() {
        fmt.Println("DB dibuka")
    }
    func closeDB() {
        fmt.Println("DB ditutup")
    }

    func main() {
        defer closeDB()
        openDB()

        fmt.Println("lorem ipsum")
        fmt.Println("dolor sit amet")
    }
    ```

<br/>

## Video – Pointer (12:06)

**Pointer** adalah sebuah variabel yang menyimpan alamat memori variabel lain


<br/>
<p align="center">
<a href="https://www.instagram.com/p/CpHBRFWy6VD/?utm_source=ig_web_copy_link">
<img title="flowchart-symbol" alt="flowchart-symbol"  width= "500px" src="https://lh4.googleusercontent.com/fqOmgQJlMAHQqi2Qt7vE4sXqYTslzoE5N74KUb6A-QO6fRUupkT5pQPTJ2jJWlse6ms=w2400">
</a>
</p>

<br/>

## Video – Struct and Interfaces (13:31)

**Struct** adalah alternatif oop di Golang. Karena golang tidak mendukung oop

- Variabel object pointer
    ```golang
    type Student struct{
        Name string
        Age int
        Jurusan string
    }

    func main() {	
        person1 := Student{}
        person1.Name="joko"

        person2 := &person1

        fmt.Println("person1.Name : ", person1.Name)
        fmt.Println("person2.Name : ", person2.Name)
        
        person2.Name="eko"

        fmt.Println("person1.Name : ", person1.Name)
        fmt.Println("person2.Name : ", person2.Name)
    }
    ```

- Embeded struct
    ```golang
    type Student struct {
        Name string
        Age  int
        University
    }

    type University struct {
        Nim      string
        Jurusan  string
        Semester int
    }

    func main() {
        person1 := Student{}
        person1.Name = "joko"
        person1.Nim = "A11.2020.12121"

        fmt.Println(person1.Name)				// direct access
        fmt.Println(person1.University.Nim)		// access the parent first
    }
    ```

- Anonymous struct
    ```golang
    func main() {
        person1 := struct {
            grade int
        }{grade: 12}

        fmt.Println(person1.grade)
    }
    ```

- Kombinasi struct dan slice 
    ```golang
    type Student struct {
        name string
        age  int
    }

    func main() {
        allStudent := []Student{
            {name: "joko", age: 22},
            {name: "eko", age: 21},
            {name: "bejo", age: 23},
        }

        fmt.Println(allStudent) // [{joko 22} {eko 21} {bejo 23}]
        for _, item := range allStudent {
            fmt.Println(item)
            // {joko 22}
            // {eko 21}
            // {bejo 23}
        }
    }
    ```

<br/>

## Video – Methods (8:47)

**Methods** adalah sebuah fungsi yang di masukan pada sebuah tipe data. Misalnya struct. Keunggulan method dibanding fungsi biasa adalah memiliki akses ke property struct hingga level ***private***

```golang
type Student struct {
	name string
	age  int
}

func (s Student) SayHello() {
	fmt.Println("Hello", s.name)
}

func (s Student) SayAge() {
	fmt.Println("Your age", s.age)
}

func main() {
	Person1 := Student{"joko", 22}
	Person1.SayHello()
	Person1.SayAge()
}
```

- Method pointer
    ```golang
    type Student struct {
        name string
        age  int
    }

    func (s Student) ChangeName1(name string) {
        s.name = name
    }

    func (s *Student) ChangeName2(name string) {
        s.name = name
    }

    func main() {
        Person1 := Student{"joko", 22}
        fmt.Println(Person1.name) 	//joko

        Person1.ChangeName1("eko") //non pointer method
        fmt.Println(Person1.name)  //joko

        Person1.ChangeName2("eko") //pointer method
        fmt.Println(Person1.name)  //eko
    }
    ```

<br/>

## Errors, panic and recover keywords

- Errors
    ```golang
    import (
        "errors"
        "fmt"
    )

    func myFunc(i int) (int, error) {
        if i <= 0 {
            return i, errors.New("should be greater than zero")
        }
        return i, nil
    }

    func main() {
        fmt.Println(myFunc(1))
        fmt.Println(myFunc(-2))
    }
    ```

- Panic
    ```golang
    import (
        "fmt"
        "strconv"
    )

    func main() {
        var input string
        fmt.Print("Type number: ")
        fmt.Scanln(&input)

        number, err := strconv.Atoi(input)

        if err == nil {
            fmt.Println(number, "is number")
        } else {
            fmt.Println(input, "is not number")
            panic(err.Error())
        }
    }
    ```

- Recover
    ```golang
    import (
        "errors"
        "fmt"
        "strings"
    )

    func validate(input string) (bool, error) {
        if strings.TrimSpace(input) == "" {
            return false, errors.New("cannot be empty")
        }
        return true, nil
    }

    func catch() {
        if r := recover(); r != nil {
            fmt.Println("Error occured:", r)
        } else {
            fmt.Println("Applicatipoon running perfecto")
        }
    }

    func main() {
        defer catch()

        var name string
        fmt.Print("Type name: ")
        fmt.Scanln(&name)

        if valid, err := validate(name); valid {
            fmt.Println("halo", name)
        } else {
            panic(err.Error())
            fmt.Println("end")
        }
    }
    ```

> di Golang kita bisa membuat custom data types

> method dapat diasosiakan dengan tipe data tertentu dan dapat dipanggil pada variabel tipe data tersebut