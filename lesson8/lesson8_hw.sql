-- Active: 1712133551411@@127.0.0.1@5432@techco@public

create table customers (
	id INT PRIMARY KEY,
	first_name VARCHAR(50),
	last_name VARCHAR(50)
);

create table products (
	id INT PRIMARY KEY,
	product_name VARCHAR(50)
);

create table orders (
	id INT PRIMARY KEY,
	product_id INT,
	customer_id INT,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);

create table order_item (
	id INT,
	order_id INT,
	order_item VARCHAR(11),
    FOREIGN KEY (order_id) REFERENCES orders(id)
);

CREATE INDEX idx_customer ON customers(id);
CREATE INDEX idx_products ON products(id);
CREATE INDEX idx_orders ON orders(id);
CREATE INDEX idx_item ON order_item(order_item);


SELECT
    c.first_name AS customer_name,
    p.product_name,
    COUNT(p.product_name) AS product_count,
    oi.order_item
FROM
    customers c
JOIN
    orders o ON c.id = o.customer_id
JOIN
    products p ON o.product_id = p.id
JOIN
    order_item oi ON o.id = oi.order_id
GROUP BY
    c.first_name, p.product_name, oi.order_item;


SELECT * FROM customers;
SELECT * FROM orders;
SELECT * FROM products;
SELECT * FROM order_item;