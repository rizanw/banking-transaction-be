package module

import (
	"context"
	"errors"
	duser "tx-bank/internal/domain/user"

	"github.com/google/uuid"
)

var errUserNotFound = errors.New("user not found")

// corporateMakers returns the requester and the IDs of every user in the
// requester's corporate. Transactions are only visible within a corporate.
func (u *usecase) corporateMakers(ctx context.Context, requesterID uuid.UUID) (*duser.User, []uuid.UUID, error) {
	requester, err := u.userRepo.GetUserByID(ctx, requesterID)
	if err != nil {
		return nil, nil, err
	}
	if requester == nil {
		return nil, nil, errUserNotFound
	}

	users, err := u.userRepo.FindUsers(ctx, nil, nil, &requester.Corporate.ID)
	if err != nil {
		return nil, nil, err
	}
	makers := make([]uuid.UUID, 0, len(users))
	for _, user := range users {
		makers = append(makers, user.ID)
	}
	return requester, makers, nil
}

// sameCorporate reports whether the transaction's maker belongs to the requester's corporate.
func (u *usecase) sameCorporate(ctx context.Context, requesterID, makerID uuid.UUID) (bool, error) {
	requester, err := u.userRepo.GetUserByID(ctx, requesterID)
	if err != nil || requester == nil {
		return false, err
	}
	maker, err := u.userRepo.GetUserByID(ctx, makerID)
	if err != nil || maker == nil {
		return false, err
	}
	return requester.Corporate.ID == maker.Corporate.ID, nil
}
