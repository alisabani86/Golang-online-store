ALTER TABLE products
ADD COLUMN quantity INT NOT NULL;

ALTER TABLE shoping_cart
ADD COLUMN created_at TIMESTAMP;