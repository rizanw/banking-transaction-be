-- create users table
CREATE TABLE IF NOT EXISTS "users"
(
    "id"           bigserial PRIMARY KEY,
    "username"     varchar(255) NOT NULL UNIQUE,
    "password"     varchar(255) NOT NULL,
    "email"        varchar(255) NOT NULL UNIQUE,
    "phone"        varchar(20),
    "corporate_id" bigint       NOT NULL,
    "role"         smallint     NOT NULL,
    "created_at"   timestamptz  NOT NULL DEFAULT (now()),
    "updated_at"   timestamptz
);

CREATE UNIQUE INDEX CONCURRENTLY ON users ("username");

ALTER TABLE "users"
    ADD FOREIGN KEY ("corporate_id") REFERENCES "corporates" ("id");

/*DROP TABLE users; */
