-- create transactions table
CREATE TABLE "transactions"
(
    "id"               UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    "ref_num"          varchar(255) NOT NULL UNIQUE,
    "amount_total"     numeric      NOT NULL,
    "record_total"     int          NOT NULL,
    "maker"            UUID,
    "date"             timestamptz  NOT NULL,
    "status"           smallint     NOT NULL,
    "instruction_type" varchar(255) NOT NULL,
    "created_at"       timestamptz  NOT NULL DEFAULT (now()),
    "updated_at"       timestamptz  NOT NULL DEFAULT (now())
);

ALTER TABLE "transactions"
    ADD FOREIGN KEY ("maker") REFERENCES "users" ("id");

/*DROP TABLE corporates; */
