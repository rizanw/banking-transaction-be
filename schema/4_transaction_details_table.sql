-- create transaction_details table
CREATE TABLE IF NOT EXISTS "transaction_details"
(
    "id"              bigserial PRIMARY KEY,
    "transaction_id"  bigint,
    "to_account_num"  varchar(255) NOT NULL,
    "to_account_name" varchar(255) NOT NULL,
    "to_account_bank" varchar(255) NOT NULL,
    "amount"          numeric      NOT NULL,
    "description"     text,
    "status"          smallint     NOT NULL,
    "created_at"      timestamptz  NOT NULL DEFAULT (now()),
    "updated_at"      timestamptz
);

ALTER TABLE "transaction_details"
    ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

/*DROP TABLE transaction_details; */
