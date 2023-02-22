# Section 5 - Basic Programming
## Video – Introduction Golang (11:10)

Apa itu **Golang** :
- Bahasa pemrograman khusus BE
- Open Source
- Lightweight

Workspace Golang ada 2:
1. di direktori %GOPATH%, terdapat 3 folder :
   - pkg
   - bin
   - src : disinilah tempat ngodingnya
2. dimanapun, tpi jangan lupa setel :
   ```golang
   go mod init example/example
   ```

Beberapa aturan penting di Golang :
- hanya ada satu func main dalam satu package
- untuk memanggil func dari beda file tpi satu package, haru menggunakan :
    ```golang
    go run . 
    // atau
    go build main.go
    ```



