-- create corporates table
CREATE TABLE "corporates"
(
    "id"          UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    "account_num" varchar(255) NOT NULL UNIQUE,
    "name"        varchar(255) NOT NULL,
    "created_at"  timestamptz  NOT NULL DEFAULT (now()),
    "updated_at"  timestamptz  NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX CONCURRENTLY ON corporates ("account_num");

/*DROP TABLE corporates; */
