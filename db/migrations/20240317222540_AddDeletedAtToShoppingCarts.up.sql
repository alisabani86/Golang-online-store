-- Up migration
ALTER TABLE shoping_cart
ADD COLUMN deleted_at TIMESTAMP;

