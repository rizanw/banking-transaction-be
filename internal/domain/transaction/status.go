package transaction

type Status int

const (
	StatusNone             Status = iota
	StatusAwaitingApproval Status = iota
	StatusApproved         Status = iota
	StatusRejected         Status = iota
)

func (s Status) GetName() string {
	switch s {
	case StatusAwaitingApproval:
		return "Awaiting Approval"
	case StatusApproved:
		return "Approved"
	case StatusRejected:
		return "Rejected"
	}
	return ""
}

type StatusCounter struct {
	AwaitingApproval int64
	StatusApproved   int64
	StatusRejected   int64
}
