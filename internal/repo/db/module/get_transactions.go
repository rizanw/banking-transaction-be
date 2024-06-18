package module

import (
	"fmt"
	"strings"
	"tx-bank/internal/model/transaction"

	"github.com/lib/pq"
)

func (r *repo) GetTransactions(filter transaction.TransactionFilter, offset, limit int) ([]transaction.TransactionDB, int32, error) {
	var (
		results []transaction.TransactionDB
		query   string
		err     error
	)

	query, whereList, whereParam, params := buildQueryGetTransactions(filter, offset, limit)
	rows, err := r.db.Query(query, params...)
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		var tx transaction.TransactionDB
		err = rows.Scan(
			&tx.ID,
			&tx.RefNum,
			&tx.AmountTotal,
			&tx.RecordTotal,
			&tx.Maker,
			&tx.TxDate,
			&tx.Status,
		)
		if err != nil {
			return nil, 0, err
		}
		results = append(results, tx)
	}

	var total int32
	params = []interface{}{}
	query = "SELECT COUNT(*) FROM transactions"
	if len(whereList) > 0 {
		query = query + " WHERE " + strings.Join(whereList, " AND ")
	}
	err = r.db.QueryRow(query, whereParam...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return results, total, err
}

func buildQueryGetTransactions(filter transaction.TransactionFilter, offset, limit int) (string, []string, []interface{}, []interface{}) {
	var (
		query       string = qGetTransactions
		whereList   []string
		whereParams []interface{}
		params      []interface{}
		countParams int = 1
	)

	if filter.Status > 0 {
		whereList = append(whereList, fmt.Sprintf("status=$%d", countParams))
		params = append(params, filter.Status)
		countParams++
	}
	if len(filter.Makers) > 0 {
		whereList = append(whereList, fmt.Sprintf(`maker=ANY($%d)`, countParams))
		params = append(params, pq.Array(filter.Makers))
		countParams++
	}
	if !filter.StartDate.IsZero() {
		whereList = append(whereList, fmt.Sprintf("date>=$%d", countParams))
		params = append(params, filter.StartDate)
		countParams++
	}
	if !filter.EndDate.IsZero() {
		whereList = append(whereList, fmt.Sprintf("date<$%d", countParams))
		params = append(params, filter.EndDate)
		countParams++
	}
	if len(whereList) > 0 {
		query = query + " WHERE " + strings.Join(whereList, " AND ")
		whereParams = append(whereParams, params...)
	}

	query = query + " ORDER BY date DESC"
	if offset >= 0 {
		query = query + fmt.Sprintf(" OFFSET $%d", countParams)
		params = append(params, offset)
		countParams++
	}
	if limit > 0 {
		query = query + fmt.Sprintf(" LIMIT $%d", countParams)
		params = append(params, limit)
		countParams++
	}

	return query, whereList, whereParams, params
}
