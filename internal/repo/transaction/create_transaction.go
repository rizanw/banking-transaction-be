package transaction

import (
	"context"
	domain "tx-bank/internal/domain/transaction"
)

const (
	qCreateTransaction = `
	INSERT INTO "transactions"
		("id","ref_num", "amount_total", "record_total", "maker", "date", "status", "instruction_type")
	VALUES 
		($1,$2,$3,$4,$5,$6,$7,$8);
`
	qCreateTransactionDetail = `
	INSERT INTO "transaction_details"
		("id", "transaction_id", "to_account_num", "to_account_name", "to_account_bank", "amount", "description", "status")
	VALUES
		($1,$2,$3,$4,$5,$6,$7,$8);
`
)

// CreateTransaction stores the transaction and all its details in one database transaction.
func (r *repo) CreateTransaction(ctx context.Context, trx *domain.Transaction) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() //nolint:errcheck // no-op after commit

	if _, err = tx.ExecContext(ctx,
		qCreateTransaction,
		trx.ID,
		trx.RefNum,
		trx.AmountTotal,
		trx.RecordTotal,
		trx.Maker.ID,
		trx.TxDate,
		trx.Status,
		trx.InstructionType,
	); err != nil {
		return err
	}

	for _, d := range trx.Details {
		if _, err = tx.ExecContext(ctx,
			qCreateTransactionDetail,
			d.ID,
			trx.ID,
			d.ToAccountNumber,
			d.ToAccountName,
			d.ToAccountBank,
			d.Amount,
			d.Description,
			d.Status,
		); err != nil {
			return err
		}
	}

	return tx.Commit()
}
