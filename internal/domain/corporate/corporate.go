package corporate

import (
	"github.com/google/uuid"
)

type Corporate struct {
	ID         uuid.UUID
	Name       string
	AccountNum string
}
