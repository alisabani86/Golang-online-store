


CREATE TABLE "accounts"(
    "id" INT PRIMARY KEY NOT NULL,
    "user_id" INT  NOT NULL,
    "account_number" VARCHAR,
    "balance" INT

);

CREATE TABLE "transactions" (
    "id" INT PRIMARY KEY NOT NULL,
    "sof_number" VARCHAR,
    "dof_number" VARCHAR,
    "account_id" INT NOT NULL,
    "amount" DECIMAL NOT NULL,
    "transaction_type" INT,
    "transactions_datetime" timestamp
)