package module

import (
	"fmt"
	"tx-bank/internal/common/session"
)

func (u *usecase) Logout(sessionKey string) error {
	if _, exists := session.Sessions[sessionKey]; !exists {
		return fmt.Errorf("invalid token")
	}

	delete(session.Sessions, sessionKey)
	return nil
}
