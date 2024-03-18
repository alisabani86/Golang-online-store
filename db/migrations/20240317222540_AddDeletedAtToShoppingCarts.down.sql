-- Down migration (to revert changes)
ALTER TABLE shoping_cart
DROP COLUMN deleted_at;

