package module

const qFindCorporate = `
		SELECT 
			"id",
			"name",
			"account_num"
		FROM
		    "corporates"
		WHERE
		    account_num=$1;
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
	WHERE username = $1 OR email = $2;
`
