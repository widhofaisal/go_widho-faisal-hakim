# Section 16 - Join - Union - Agregasi - Subquery - Function (DBMS)

>Database Manipulation Language (DML) adalah bahasa pemrograman yang digunakan untuk mengubah, menambah, atau menghapus data dari sebuah database. DML terdiri dari beberapa perintah seperti SELECT, INSERT, UPDATE, DELETE, dan lain-lain. Perintah-perintah ini digunakan untuk memanipulasi data dalam sebuah tabel atau relasi.

ex:
```sql
-- SELECT statement untuk mengambil data dari tabel
SELECT * FROM customers;

-- INSERT statement untuk menambah data ke dalam tabel
INSERT INTO customers (name, email, phone) VALUES ('John Doe', 'john.doe@example.com', '123-456-789');

-- UPDATE statement untuk mengubah data pada tabel
UPDATE customers SET phone = '987-654-321' WHERE name = 'John Doe';

-- DELETE statement untuk menghapus data dari tabel
DELETE FROM customers WHERE id = 1;

```
</br>

>JOIN adalah salah satu operasi dalam SQL yang digunakan untuk menggabungkan dua atau lebih tabel berdasarkan kolom atau atribut yang sama. JOIN digunakan ketika kita ingin mengambil data dari dua tabel atau lebih yang berelasi satu sama lain. Ada beberapa jenis JOIN, seperti INNER JOIN, LEFT JOIN, RIGHT JOIN, dan FULL JOIN, yang masing-masing digunakan untuk tujuan yang berbeda.
ex:
```sql
-- INNER JOIN untuk menggabungkan data dari dua tabel
SELECT orders.id, customers.name, orders.order_date, order_items.product_name, order_items.quantity
FROM orders
INNER JOIN customers ON orders.customer_id = customers.id
INNER JOIN order_items ON orders.id = order_items.order_id;

-- LEFT JOIN untuk mengambil semua data dari tabel kiri dan data yang sesuai dari tabel kanan
SELECT customers.name, orders.order_date, order_items.product_name, order_items.quantity
FROM customers
LEFT JOIN orders ON customers.id = orders.customer_id
LEFT JOIN order_items ON orders.id = order_items.order_id;

-- RIGHT JOIN untuk mengambil semua data dari tabel kanan dan data yang sesuai dari tabel kiri
SELECT orders.order_date, order_items.product_name, order_items.quantity, customers.name
FROM orders
RIGHT JOIN order_items ON orders.id = order_items.order_id
RIGHT JOIN customers ON orders.customer_id = customers.id;
```
</br>

>UNION adalah operasi dalam SQL yang digunakan untuk menggabungkan hasil dari dua SELECT statements atau lebih menjadi satu set data. UNION menggabungkan data dari dua atau lebih tabel atau hasil query yang memiliki struktur yang sama. UNION bisa digunakan untuk menggabungkan hasil query dari tabel atau relasi yang berbeda untuk tujuan analisis atau pelaporan.
```sql
-- UNION untuk menggabungkan hasil dari dua SELECT statements
SELECT product_name, price FROM products WHERE category = 'Electronic'
UNION
SELECT product_name, price FROM products WHERE category = 'Fashion';
```
</br>

>Dalam penggunaan database, DML, JOIN, dan UNION merupakan materi yang sangat penting dan sering digunakan dalam pemrograman SQL. Dengan menggunakan perintah DML, kita dapat memanipulasi data dalam tabel, seperti mengambil data dari tabel, menambah, mengubah, atau menghapus data dari tabel. Dalam operasi JOIN, kita dapat menggabungkan data dari dua tabel atau lebih yang berelasi, sedangkan dengan UNION kita dapat menggabungkan data dari dua query atau tabel yang memiliki struktur yang sama. Dengan memahami konsep dasar dan penggunaan DML, JOIN, dan UNION, kita dapat membuat query SQL yang lebih kompleks dan efisien untuk memanipulasi data dalam database.