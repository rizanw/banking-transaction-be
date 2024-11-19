package transaction

import (
	"context"
	"fmt"
	"strings"
	domain "tx-bank/internal/domain/transaction"

	"github.com/lib/pq"
	"golang.org/x/sync/errgroup"
)

const qGetTransactions = `
	SELECT 
		"id", "ref_num", "amount_total", "record_total", "maker", "date", "status"
	FROM 
	    "transactions"
`

func (r *repo) GetTransactions(ctx context.Context, in domain.TransactionFilter) (int32, []domain.Transaction, error) {
	var (
		transactions []domain.Transaction
		total        int32
		eg           errgroup.Group
	)

	query, whereList, whereParam, params := buildQueryGetTransactions(in, in.Offset, in.Limit)

	eg.Go(func() error {
		rows, err := r.db.QueryxContext(ctx, query, params...)
		if err != nil {
			return err
		}
		for rows.Next() {
			var tx domain.Transaction
			if err = rows.Scan(
				&tx.ID,
				&tx.RefNum,
				&tx.AmountTotal,
				&tx.RecordTotal,
				&tx.Maker.ID,
				&tx.TxDate,
				&tx.Status,
			); err != nil {
				return err
			}

			transactions = append(transactions, tx)
		}
		return nil
	})

	eg.Go(func() error {
		query = "SELECT COUNT(*) FROM transactions"
		if len(whereList) > 0 {
			query = query + " WHERE " + strings.Join(whereList, " AND ")
		}
		if err := r.db.QueryRow(query, whereParam...).Scan(&total); err != nil {
			return err
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return 0, []domain.Transaction{}, err
	}

	return total, transactions, nil
}

func buildQueryGetTransactions(filter domain.TransactionFilter, offset, limit int) (string, []string, []interface{}, []interface{}) {
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
