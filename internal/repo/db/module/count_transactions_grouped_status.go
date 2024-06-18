package module

func (r *repo) CountTransactionsGroupedStatus() (map[int32]int64, error) {
	var (
		res = make(map[int32]int64, 0)
	)

	rows, err := r.db.Query(qCountTransactionsGroupedStatus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			status int32
			count  int64
		)
		rows.Scan(&status, &count)
		res[status] = count
	}

	return res, nil
}
