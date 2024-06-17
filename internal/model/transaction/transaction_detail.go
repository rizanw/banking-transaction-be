package transaction

type TransactionDetailDB struct {
	ID              int64   `db:"id"`
	TransactionID   int64   `db:"transaction_id"`
	ToAccountNumber string  `db:"to_account_num"`
	ToAccountName   string  `db:"to_account_name"`
	ToAccountBank   string  `db:"to_account_bank"`
	Amount          float64 `db:"amount"`
	Description     string  `db:"description"`
	Status          int32   `db:"status"`
}
