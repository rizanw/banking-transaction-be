-- create otps table
CREATE TABLE "otps"
(
    "id"         bigserial PRIMARY KEY,
    "user_id"    bigint      NOT NULL,
    "code"       varchar(6)  NOT NULL,
    "expires_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

/*DROP TABLE otps; */
