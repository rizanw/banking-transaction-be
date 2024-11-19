-- create audit_logs table
CREATE TABLE "audit_logs"
(
    "id"             UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    "transaction_id" UUID         NOT NULL,
    "user_id"        UUID         NOT NULL,
    "action"         varchar(255) NOT NULL,
    "timestamp"      timestamptz  NOT NULL DEFAULT (now())
);

ALTER TABLE "audit_logs"
    ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

ALTER TABLE "audit_logs"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

/*DROP TABLE audit_logs; */
