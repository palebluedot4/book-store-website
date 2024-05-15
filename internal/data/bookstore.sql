CREATE DATABASE bookstore;

USE bookstore;

CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT, username VARCHAR(100) NOT NULL UNIQUE, password VARCHAR(100) NOT NULL, email VARCHAR(100)
)

SELECT * FROM bookstore.users;

CREATE TABLE books (
    id INT PRIMARY KEY AUTO_INCREMENT, title VARCHAR(100) NOT NULL, author VARCHAR(100) NOT NULL, price DOUBLE(11, 2) NOT NULL, sales INT NOT NULL, stock INT NOT NULL, img_path VARCHAR(100)
)

SELECT * FROM bookstore.books;

CREATE TABLE sessions (
    session_id VARCHAR(100) PRIMARY KEY, username VARCHAR(100) NOT NULL, user_id INT NOT NULL, FOREIGN KEY (user_id) REFERENCES users (id)
)

SELECT * FROM bookstore.sessions;

CREATE TABLE carts (
    id VARCHAR(100) PRIMARY KEY, total_count INT NOT NULL, total_amount DOUBLE(11, 2) NOT NULL, user_id INT NOT NULL, FOREIGN KEY (user_id) REFERENCES users (id)
)

SELECT * FROM bookstore.carts;

CREATE TABLE cart_items (
    id INT PRIMARY KEY AUTO_INCREMENT, count INT NOT NULL, amount DOUBLE(11, 2) NOT NULL, book_id INT NOT NULL, cart_id VARCHAR(100) NOT NULL, FOREIGN KEY (book_id) REFERENCES books (id), FOREIGN KEY (cart_id) REFERENCES carts (id)
)

SELECT * FROM bookstore.cart_items;

CREATE TABLE orders (
    id VARCHAR(100) PRIMARY KEY, create_time DATETIME NOT NULL, total_count INT NOT NULL, total_amount DOUBLE(11, 2) NOT NULL, status INT NOT NULL, user_id INT, FOREIGN KEY (user_id) REFERENCES users (id)
)

SELECT * FROM bookstore.orders;

CREATE TABLE order_items (
    id INT PRIMARY KEY AUTO_INCREMENT, count INT NOT NULL, amount DOUBLE(11, 2) NOT NULL, title VARCHAR(100) NOT NULL, author VARCHAR(100) NOT NULL, price DOUBLE(11, 2) NOT NULL, img_path VARCHAR(100) NOT NULL, order_id VARCHAR(100) NOT NULL, FOREIGN KEY (order_id) REFERENCES orders (id)
)

SELECT * FROM bookstore.order_items;