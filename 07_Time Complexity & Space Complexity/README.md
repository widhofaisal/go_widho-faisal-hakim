# Section 7 : Time Complexity & Space Complexity

## Video – Time Complexity (26:31)

**Time Complexity** adalah proses yang paling dominan, dilambangkan dengan O

### Fun Fact :
> "Time complexity yang lebih cepat belum tentu cepat dalam hitungan waktu (sekon). Karena time complexity diasumsikan dari seberapa banyak proses yang paling dominan di eksekusi. Bisa jadi model A looping nya sedikit tpi dalam setiap looping itu menjalankan proses yang banyak/besar/berat. Disamping itu model B looping nya banyak namun dalam setiap looping itu hanyak menjalankan proses yang kecil/ringan. Semuanya bersifat **RELATIF**." -Widho wkwk

<br/>

Jenis-jenis time complexity
- Linear time : O(n)
    ```golang
    func sum(n int) int {
        result := 0
        for i := 0; i < n; i++ {
            result += i
        }
        return result
    }
    ```
    Time complexity dari kode di atas adalah O(n), karena loop for di dalam fungsi sum dijalankan sebanyak n kali, di mana n adalah nilai input yang diberikan ke fungsi. Dalam setiap iterasi, operasi penjumlahan dilakukan satu kali. Oleh karena itu, waktu yang dibutuhkan untuk mengeksekusi fungsi sum secara linear bergantung pada nilai input n. Oleh karena itu, kompleksitas waktu dari fungsi sum adalah O(n).

    Other example :
    ```golang
    func sum(n []int) []int {
        result := []int{}
        for i := len(n); i > 0; i-- {
            result = append(result, i+10)
        }
        return result
    }
    ```
    Time complexity dari kode di atas adalah O(n), karena loop for di dalam fungsi sum dijalankan sebanyak n kali, di mana n adalah jumlah elemen dalam slice n. Dalam setiap iterasi, fungsi append dipanggil untuk menambahkan elemen baru ke slice result, yang membutuhkan waktu yang secara linear bergantung pada ukuran slice result. Oleh karena itu, kompleksitas waktu dari fungsi sum adalah O(n).


- Constant time : O(1)
    ```golang
    func sum(n int) int {
        result := 0
        for i := 0; i < 10; i++ {
            result += 1
        }
        return result
    }
    ```
    Time complexity dari kode di atas adalah O(1), karena loop for di dalam fungsi sum hanya dijalankan sebanyak 10 kali, yaitu sebanyak nilai konstan yang telah ditentukan. Oleh karena itu, waktu yang dibutuhkan untuk mengeksekusi fungsi sum tidak tergantung pada ukuran input n, melainkan hanya bergantung pada nilai konstan yaitu 10. Sehingga, kompleksitas waktu fungsi sum adalah konstan atau O(1).

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
    func sum(n int) int {
        result := 0
        for i := n; i > 1; i /= 2 {
            result += i
        }
        return result
    }
    ```
    Time complexity dari kode di atas adalah O(log n), karena loop for di dalam fungsi sum dijalankan sebanyak logaritma basis 2 dari nilai input n. Dalam setiap iterasi, nilai i dibagi 2, sehingga iterasi berkurang setengah dari iterasi sebelumnya. Dalam kasus ini, kompleksitas waktu tidak tergantung pada nilai n secara linear, melainkan logaritmik. Oleh karena itu, kompleksitas waktu dari fungsi sum adalah O(log n).
    
    other example :
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
    func sum(n int) int {
        result := 0
        for i := n; i > 0; i-- {
            for j := 0; j < n; j++ {
                result += 1
            }
        }
        return result
    }
    ```
    Time complexity dari kode di atas adalah O(n^2), karena ada dua loop yang bersarang satu sama lain. Loop pertama dijalankan sebanyak n kali, dan di dalam setiap iterasi, loop kedua dijalankan sebanyak n kali pula. Oleh karena itu, total loop yang dijalankan sebanyak n^2 kali. Dalam setiap iterasi loop kedua, hanya ada satu operasi penjumlahan yang dilakukan. Oleh karena itu, waktu yang dibutuhkan untuk mengeksekusi fungsi sum secara kuadratik bergantung pada nilai input n. Oleh karena itu, kompleksitas waktu dari fungsi sum adalah O(n^2).

    Other example :
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

<br/>

