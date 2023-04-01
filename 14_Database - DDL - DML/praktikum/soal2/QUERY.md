# Kumpulan query untuk mengerjakan soal nomor 2
- 1: 
    ```sql
    CREATE TABLE users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255),
        address VARCHAR(255),
        birthdate DATE,
        status BOOLEAN DEFAULT '1',
        gender CHAR,
        createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updatedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        );
    ```
- 2a: 
    ```sql
    CREATE TABLE products_type(
        id INT AUTO_INCREMENT PRIMARY KEY,
        type VARCHAR(255), 
        expired TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    ```
    ```sql
    CREATE TABLE operators(
        id INT AUTO_INCREMENT PRIMARY KEY,
        company VARCHAR(255),
        address VARCHAR(255),
        phone VARCHAR(255)
        );
    ```
    ```sql
    CREATE TABLE products_desc(
        id INT AUTO_INCREMENT PRIMARY KEY,
        description VARCHAR(255),
        history VARCHAR(255)
        );
    ```
    ```sql
    CREATE TABLE products(
        id INT AUTO_INCREMENT PRIMARY KEY,
        product_name VARCHAR(255),
        type_id INT,
        operator_id INT,
        description_id INT,
        FOREIGN KEY(type_id) REFERENCES products_type(id),
        FOREIGN KEY(operator_id) REFERENCES operators(id),
        FOREIGN KEY(description_id) REFERENCES products_desc(id)
        );
    ```
    ```sql
    CREATE TABLE payment_method(
        id INT AUTO_INCREMENT PRIMARY KEY,
        method VARCHAR(255),
        interest FLOAT
        );
    ```
- 2b:
    ```sql
    CREATE TABLE transactions(
        id INT AUTO_INCREMENT PRIMARY KEY,
        buyer_id INT,
        product_id INT,
        amount INT,
        total_price FLOAT,
        payment_id INT,
        FOREIGN KEY(buyer_id) REFERENCES users(id),
        FOREIGN KEY(product_id) REFERENCES products(id),
        FOREIGN KEY(payment_id) REFERENCES payment_method(id)
        );
    ```
    ```sql
    CREATE TABLE transaction_detail(
        id INT AUTO_INCREMENT PRIMARY KEY,
        transaction_id INT,
        shipping_address VARCHAR(255),
        FOREIGN KEY(transaction_id) REFERENCES transactions(id)
        );
    ```
- 3:
    ```sql
    CREATE TABLE kurir(
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255),
        createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updatedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        );
    ```
- 4:
    ```sql
    ALTER TABLE kurir
    ADD ongkos_dasar FLOAT;
    ```
- 5:
    ```sql
    RENAME TABLE kurir TO shipping;
    ```
- 6:
    ```sql
    DROP TABLE shipping;
    ```
- 7a:
    ```sql
    CREATE TABLE payment_method_description(
        id INT AUTO_INCREMENT PRIMARY KEY,
        description VARCHAR(255),
        rating INT
        );
    ```
    ```sql
    ALTER TABLE payment_method
    ADD description_id INT;
    ```
    ```sql
    ALTER TABLE payment_method
    ADD CONSTRAINT description_id
    FOREIGN KEY(description_id) REFERENCES payment_method_description(id);
    ```
- 7b:
    ```sql
    ALTER TABLE users
    DROP address;
    ```
    ```sql
    CREATE TABLE address(
        id INT AUTO_INCREMENT PRIMARY KEY,
        address VARCHAR(255),
        difficulty INT,
        user_id INT,
        FOREIGN KEY(user_id) REFERENCES users(id)
        );
    ```
- 7c:
    ```sql
    CREATE TABLE user_payment_method_detail(
        user_id INT,
        payment_method_id INT,
        FOREIGN KEY(user_id) REFERENCES users(id),
        FOREIGN KEY(payment_method_id) REFERENCES payment_method(id)
        );
    ```