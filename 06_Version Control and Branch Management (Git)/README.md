# Section 6 - Version Control and Branch Management (Git)

## Video KMGolang – Version Control and Branch Management (Git) (1:26:51)

**Git** memudahkan user untuk melakukan tracking terhadap history perubahan file. Dalam hal ini adalah file kodingan untuk programmer bekerjasama dengan orang lain.

Repository management populer yang terintegrasi Git :
- Github
- Gitlab
- Bitbucket

Catatan penting tentang penggunaan **Github** :
- file README.md berisi panduan penggunaan aplikasi dan dokumentasi
- nama file yang dicantumkan di file .gitignore tidak akan di push ke Github (tetap berada di Local Computer)
  
<br/>

Perintah **Git** yang dipelajari :

- login git di Local Computer
    ```git
    git config --global user.name "widho faisal"                #ganti widho faisal dengan username github kamu
    git config --global user.email "alterragolang@gmail.com"    #ganti email dengan email kamu yang digunakan mendaftar github
    ```

- cek versi git
    ```git
    git --version   #akan menampilkan versi git yang terinstall
    ```

- menggunakan repository lewat 2 cara : 
  - create repository di Github, lalu clone di Local Computer
    ```git
    git clone https://github.com/mnfirdauss/git_learning.git    #akan melakukan download repository ke Local Computer
    ```
  - create repository github, lalu add remote di folder Local Computer
    ```git
    #buat perubahan seperti membuat file .txt di dalam sebuah folder
    git init
    git remote add origin https://github.com/mnfirdauss/git_learning.git
    git status                      #akan muncul daftar perubahan yang terjadi
    git add .                       #untuk menambahkan semua perubahan ke Staging Area
    git commit -m "tuliskan pesan"  #membuat commit dan memberi pesan
    git push -u origin main         #melakukan push ke remote bernama origin dan branch bernama main
    ```

- push ke Github
    ```git
    git status                      #akan muncul daftar perubahan yang terjadi
    git add .                       #untuk menambahkan semua perubahan ke Staging Area
    git commit -m "tuliskan pesan"  #membuat commit dan memberi pesan
    git push -u origin main         #melakukan push ke remote bernama origin dan branch bernama main
    ```
- mengembalikan commit yang sudah di push
  ```git
  git reset --hard 8d27235737ccc7183c48e096a0aae1d2689596d9  #ganti kode SHA sesuai commit kamu yang ingin di kembalikan
  ```

- tentang branch
  ```git
  git branch develop        #membuat branch baru bernama develop
  git checkout develop      #berpindah workingspace ke branch develop
  git branch -D develop     #menghapus branch develop
  ```