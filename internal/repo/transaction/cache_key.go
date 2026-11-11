package transaction

import "github.com/google/uuid"

func cacheKeyTransaction(id uuid.UUID) string {
	return "transaction:" + id.String()
}
