package corporate

type CorporateDB struct {
	ID         int64  `db:"id"`
	Name       string `db:"name"`
	AccountNum string `db:"account_num"`
}
