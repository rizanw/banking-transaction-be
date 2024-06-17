-- create audit_logs table
CREATE TABLE IF NOT EXISTS "audit_logs"
(
    "id"             bigserial PRIMARY KEY,
    "transaction_id" bigint       NOT NULL,
    "user_id"        bigint       NOT NULL,
    "action"         varchar(255) NOT NULL,
    "timestamp"      timestamptz  NOT NULL
);

ALTER TABLE "audit_logs"
    ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

ALTER TABLE "audit_logs"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

/*DROP TABLE audit_logs; */
