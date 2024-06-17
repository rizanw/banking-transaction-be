-- create transactions table
CREATE TABLE IF NOT EXISTS "transactions"
(
    "id"               bigserial PRIMARY KEY,
    "ref_num"          varchar(255) NOT NULL UNIQUE,
    "amount_total"     numeric      NOT NULL,
    "record_total"     int          NOT NULL,
    "maker"            bigint,
    "date"             timestamptz  NOT NULL,
    "status"           smallint     NOT NULL,
    "instruction_type" varchar(255) NOT NULL,
    "created_at"       timestamptz  NOT NULL DEFAULT (now()),
    "updated_at"       timestamptz
);

ALTER TABLE "transactions"
    ADD FOREIGN KEY ("maker") REFERENCES "users" ("id");

/*DROP TABLE corporates; */
