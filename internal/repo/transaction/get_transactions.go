package transaction

import (
	"context"
	"fmt"
	"strings"
	domain "tx-bank/internal/domain/transaction"
	"tx-bank/internal/domain/user"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/sync/errgroup"
)

const (
	qGetTransactions = `
	SELECT 
		"id", "ref_num", "amount_total", "record_total", "maker", "date", "status"
	FROM 
	    "transactions"
`
	qCountTransactions = `SELECT COUNT(*) FROM "transactions"`
)

func (r *repo) GetTransactions(ctx context.Context, in domain.TransactionFilter) (int32, []domain.Transaction, error) {
	var (
		transactions = make([]domain.Transaction, 0)
		total        int32
		eg, egCtx    = errgroup.WithContext(ctx)
	)

	where, whereParams := buildTransactionsWhere(in)

	eg.Go(func() error {
		query, params := buildQueryGetTransactions(where, whereParams, in.Offset, in.Limit)
		rows, err := r.db.QueryxContext(egCtx, query, params...)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				tx      domain.Transaction
				makerID uuid.UUID
			)
			if err = rows.Scan(
				&tx.ID,
				&tx.RefNum,
				&tx.AmountTotal,
				&tx.RecordTotal,
				&makerID,
				&tx.TxDate,
				&tx.Status,
			); err != nil {
				return err
			}
			tx.Maker = &user.User{ID: makerID}
			transactions = append(transactions, tx)
		}
		return rows.Err()
	})

	eg.Go(func() error {
		return r.db.QueryRowContext(egCtx, qCountTransactions+where, whereParams...).Scan(&total)
	})

	if err := eg.Wait(); err != nil {
		return 0, []domain.Transaction{}, err
	}

	return total, transactions, nil
}

// buildTransactionsWhere returns the WHERE clause (with a leading space) and its parameters.
func buildTransactionsWhere(filter domain.TransactionFilter) (string, []interface{}) {
	var (
		conditions []string
		params     []interface{}
	)
	add := func(condition string, value interface{}) {
		params = append(params, value)
		conditions = append(conditions, fmt.Sprintf(condition, len(params)))
	}

	// always scope to the given makers, so callers never see other corporates' data
	add(`maker=ANY($%d)`, pq.Array(filter.Makers))
	if filter.Status > 0 {
		add("status=$%d", filter.Status)
	}
	if !filter.StartDate.IsZero() {
		add("date>=$%d", filter.StartDate)
	}
	if !filter.EndDate.IsZero() {
		add("date<$%d", filter.EndDate)
	}

	return " WHERE " + strings.Join(conditions, " AND "), params
}

func buildQueryGetTransactions(where string, whereParams []interface{}, offset, limit int) (string, []interface{}) {
	query := qGetTransactions + where + " ORDER BY date DESC"
	params := append([]interface{}{}, whereParams...)

	if offset > 0 {
		params = append(params, offset)
		query += fmt.Sprintf(" OFFSET $%d", len(params))
	}
	if limit > 0 {
		params = append(params, limit)
		query += fmt.Sprintf(" LIMIT $%d", len(params))
	}

	return query, params
}
