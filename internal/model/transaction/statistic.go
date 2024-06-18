package transaction

type TransactionStats struct {
	AwaitingApproval int64 `json:"awaiting_approval"`
	Approved         int64 `json:"approved"`
	Rejected         int64 `json:"rejected"`
}
