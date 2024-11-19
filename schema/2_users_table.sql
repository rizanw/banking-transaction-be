-- create users table
CREATE TABLE "users"
(
    "id"           UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    "username"     varchar(255) NOT NULL UNIQUE,
    "password"     varchar(255) NOT NULL,
    "email"        varchar(255) NOT NULL UNIQUE,
    "phone"        varchar(20),
    "corporate_id" UUID         NOT NULL,
    "role"         smallint     NOT NULL,
    "created_at"   timestamptz  NOT NULL DEFAULT (now()),
    "updated_at"   timestamptz  NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX CONCURRENTLY ON users ("id");
CREATE UNIQUE INDEX CONCURRENTLY ON users ("username");

ALTER TABLE "users"
    ADD FOREIGN KEY ("corporate_id") REFERENCES "corporates" ("id");

/*DROP TABLE users; */
