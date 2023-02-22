# Section 7 : Time Complexity & Space Complexity

## Video – Time Complexity (26:31)

**Time Complexity** adalah proses yang paling dominan, dilambangkan dengan O

Jenis-jenis time complexity
- Linear time : O(n)
    ```golang
    func dominant(n int){
        for i:=0; i=n; i++{
            fmt.Println("Hello Golang")
        }
    }
    ```

- Constant time : O(1)
    ```golang
    func dominant(n int){
        fmt.Println("Umur saya adalah ", n)
    }
    ```

- Linear time : O(n+m)
    ```golang
    func dominant(int n, int m){
        for i:=0; i=n; i++{
            fmt.Println("Hello Golang")
        }
        for i:=0; i=m; i++{
            fmt.Println("Hello Alterra")
        }
    }
    ```

- Logarithmic time : O(log n)
    ```golang
    func logaritmic(n int) int {
        var result int = 0
        for n>1 {
            n/=2
            result+=1
        }
        return result
    }
    ```

- Quadratic time : O(n^2)
    ```golang
    func quadratic(n int) int {
        var result int = 0
        for i:=0; i<n; i++{
            for j:=0; j<n; j++{
                result+=1
            }
        }
        return result
    }
    ```
- Exponensial : O(2^n)
- Factorial : O(n!)
  
<br/>

Cara menghitung waktu proses berjalan (untuk research Time Complexity) :
- dalam satu func main saja :
    ```golang
    func main(){
        start := time.Now()
        time.Sleep(time.Second * 2)
        
        //something doing here
        
        elapsed := time.Since(start)
        fmt.Printf("process took %s", elapsed)
    }
    ```
- lebih ringkas, dipisah di func lain :

```golang
    start := time.Now()
	defer functions.Timer(start)

	time.Sleep(time.Second * 1)
```

