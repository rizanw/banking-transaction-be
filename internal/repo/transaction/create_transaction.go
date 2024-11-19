package transaction

import (
	"context"
	domain "tx-bank/internal/domain/transaction"
)

const qCreateTransaction = `
	INSERT INTO "transactions"
		("id","ref_num", "amount_total", "record_total", "maker", "date", "status", "instruction_type")
	VALUES 
		($1,$2,$3,$4,$5,$6,$7,$8);
`

func (r *repo) CreateTransaction(ctx context.Context, trx *domain.Transaction) error {
	_, err := r.db.ExecContext(ctx,
		qCreateTransaction,
		trx.ID,
		trx.RefNum,
		trx.AmountTotal,
		trx.RecordTotal,
		trx.Maker.ID,
		trx.TxDate,
		trx.Status,
		trx.InstructionType,
	)
	if err != nil {
		return err
	}

	return nil
}
