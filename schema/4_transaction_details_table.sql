-- create transaction_details table
CREATE TABLE "transaction_details"
(
    "id"              UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    "transaction_id"  UUID,
    "to_account_num"  varchar(255) NOT NULL,
    "to_account_name" varchar(255) NOT NULL,
    "to_account_bank" varchar(255) NOT NULL,
    "amount"          numeric      NOT NULL,
    "description"     text,
    "status"          smallint     NOT NULL,
    "created_at"      timestamptz  NOT NULL DEFAULT (now()),
    "updated_at"      timestamptz  NOT NULL DEFAULT (now())
);

ALTER TABLE "transaction_details"
    ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

/*DROP TABLE transaction_details; */
