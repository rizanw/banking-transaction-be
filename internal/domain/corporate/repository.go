package corporate

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetCorporates(ctx context.Context) ([]Corporate, error)
	GetCorporateByID(ctx context.Context, id uuid.UUID) (*Corporate, error)
	FindCorporate(ctx context.Context, id uuid.UUID, accountNum string) (*Corporate, error)
}
