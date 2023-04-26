# Section 22 - Middleware

> ada banyak third party untuk menerapkan middleware di golang, kali ini membahas penerapan menggunakan echo middleware (library milik echo)

### Echo Middleware, dibagi 2:
1. Echo#Pre() : adalah middleware yang dijalankan sebelum mengakses API
2. Echo#Use() : adalah middleware yang dijalankan sesudah mengakses API

### Jenis jenis middleware:
- Logger, berfungsi menampilkan output di cmd ketika client mengakses API yaitu berupa
  ```
  {"time":"2017-01-12T08:58:07.372015644-08:00","remote_ip":"::1","host":"localhost:1323","method":"GET","uri":"/","status":200,"error":"","latency":14743,"latency_human":"14.743µs","bytes_in":0,"bytes_out":2}
  ```
- Auth Basic, berfungsi menolak atau memberi akses kepada klien ketika mengakses API
  - cara kerjanya adalah melakukan pengecekan terhadap headers dari client (base64 berupa kombinasi misal email:password) apakah cocok dengan email dan password di database
- JWT