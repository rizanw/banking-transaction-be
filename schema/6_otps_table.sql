-- create otps table
CREATE TABLE "otps"
(
    "id"         UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    "recipient"  varchar(255) NOT NULL,
    "code"       varchar(6)   NOT NULL,
    "is_active"  boolean      NOT NULL DEFAULT true,
    "created_at" timestamptz  NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX CONCURRENTLY ON otps ("recipient", "code");

/*DROP TABLE otps; */
