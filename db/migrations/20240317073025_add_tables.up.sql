

CREATE TABLE "products" (
   "id" INT PRIMARY KEY NOT NULL,
   "product_name" VARCHAR NOT NULL,
   "category" VARCHAR NOT NULL,
   "price" DECIMAL NOT NULL
);

CREATE TABLE "shoping_cart" (
   "id" INT PRIMARY KEY,
   "user_id" INT NOT NULL,
   "product_id" INT NOT NULL,
   "quantity" INT NOT NULL,
   FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
   FOREIGN KEY ("product_id") REFERENCES "products" ("id")
);




CREATE TABLE "orders" (

    "id" INT PRIMARY KEY NOT NULL,
    "order_id" INT NOT NULL,
    "user_id" INT NOT NULL,
    "order_date" DATE NOT NULL,
    "total_amount" DECIMAL NOT NULL

);

CREATE TABLE "shoping_order" (
    "id" INT PRIMARY KEY,
    "order_id" INT NOT NULL,
    "cart_id" INT NOT NULL,
    FOREIGN KEY ("order_id") REFERENCES "orders" ("id"),
    FOREIGN KEY ("cart_id") REFERENCES "shoping_cart"("id")

);

