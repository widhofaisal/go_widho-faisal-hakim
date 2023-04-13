# Section 26 - Docker

## Langkah-langkah menggunakan docker:
1. Membuat Dockerfile dan mengintegrasikan dengan codingan aplikasi kamu
   ```Dockerfile
    FROM golang:1.17-alpine

    WORKDIR /app

    COPY go.mod ./
    COPY go.sum ./
    RUN go mod download

    COPY . .

    RUN go build -o /dist

    EXPOSE 3222

    CMD ["/dist"]
   ```

2. Build docker image
    ```docker
    docker build -t go-echo-gorm-mysql:v1.0.0 .
    ```
    akan menghasilkan docker image seperti ini:

    <img src="assets/build%20docker.png">

    > pentingnya menambahkan tag ke dalam docker image adalah agar kita bisa kembali ke versi sebelumnya, karena kalo tidak diberi tag maka akan merujuk default menjadi latest, dan latest ini akan menimpa/overwrite versi sebelumnya sehingga versi sebelumnya akan hilang

3. Login docker
   > kalo pertama kali biasanya diminta login dulu
    ```cmd
    docker login -u usernamedockerkamu
    ```

4. Push image ke docker hub
   - buat repository dulu di docker hub
   - buat tagging dulu
     ```cmd
     docker tag go-echo-gorm-mysql widhofaisal/go-echo-gorm-mysql
     ```
   - push 
     ```cmd
     docker push widhofaisal/go-echo-gorm-mysql
     ```