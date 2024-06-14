-- create corporates table
CREATE TABLE IF NOT EXISTS "corporates"
(
    "id"          bigserial PRIMARY KEY,
    "account_num" varchar(255) NOT NULL UNIQUE,
    "name"        varchar(255) NOT NULL,
    "created_at"  timestamptz  NOT NULL DEFAULT (now()),
    "updated_at"  timestamptz
);

CREATE UNIQUE INDEX CONCURRENTLY ON corporates ("account_num");

/*DROP TABLE corporates; */
