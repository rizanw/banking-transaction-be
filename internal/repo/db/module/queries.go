package module

const qGetCorporates = `
		SELECT 
			"id",
			"name",
			"account_num"
		FROM
		    "corporates";
`

const qFindCorporate = `
		SELECT 
			"id",
			"name",
			"account_num"
		FROM
		    "corporates"
		WHERE
		    account_num=$1 OR id=$2;
`

const qInsertCorporate = `
		INSERT INTO "corporates"
			("account_num", "name")
		VALUES 
		    ($1,$2)
		RETURNING id;
`

const qInsertUser = `
	INSERT INTO "users"
		("username", "password", "email", "phone", "corporate_id", "role")
	VALUES
		($1,$2,$3,$4,$5,$6)
	RETURNING id;
`

const qFindUser = `
	SELECT 
	    "id",
		"username",
		"password",
		"email",
		"phone",
		"corporate_id",
		"role"
	FROM "users" 
	WHERE username = $1 OR email = $2 OR id = $3 OR corporate_id = $4;
`

const qInsertTransaction = `
	INSERT INTO "transactions"
		("ref_num", "amount_total", "record_total", "maker", "date", "status", "instruction_type")
	VALUES 
		($1,$2,$3,$4,$5,$6,$7)
	RETURNING id;
`

const qGetTransactions = `
	SELECT 
		"id", "ref_num", "amount_total", "record_total", "maker", "date", "status"
	FROM 
	    "transactions"
`

const qFindTransaction = `
	SELECT 
		"id", "ref_num", "amount_total", "record_total", "maker", "date", "status", "instruction_type", "created_at"
	FROM 
	    "transactions"
	WHERE
	    "id" = $1;
`

const qUpdateTransaction = `
	UPDATE "transactions"
	SET "status" = $1, "updated_at" = $2
	WHERE "id" = $3
`

const qCountTransactionsGroupedStatus = `
	SELECT status, COUNT(*) as count
	FROM transactions
	GROUP BY status;
`

const (
	qInsertTransactionDetails = `
		INSERT INTO "transaction_details"
			("transaction_id", "to_account_num", "to_account_name", "to_account_bank", "amount", "description", "status")
		VALUES
	`
	qInsertTransactionDetailsValues = "($%d, $%d, $%d, $%d, $%d, $%d, $%d),"
)

const qFindTransactionDetails = `
	SELECT 
		"id", "transaction_id", "to_account_num", "to_account_name", "to_account_bank", "amount", "description", "status"
	FROM 
	    "transaction_details"
	WHERE
	    "transaction_id" = $1;
`

const qUpdateTransactionDetailStatus = `
	UPDATE "transaction_details" 
	SET "status" = $1, "updated_at" = $2 
	WHERE "transaction_id" = $3;
`

const qInsertAuditLog = `
	INSERT INTO "audit_logs"
		("transaction_id", "user_id", "action", "timestamp")
	VALUES 
		($1,$2,$3,$4)
	RETURNING id;
`

const qInsertOTP = `
	INSERT INTO "otps"
		("email", "code", "expires_at")
	VALUES 
		($1,$2,$3);
`

const qFindOTP = `
	SELECT 
	    "id", "email", "code", "expires_at"
	FROM "otps"
	WHERE
	    email = $1 AND code = $2;
`
