package store

import (
	"context"
)

type Store interface {
	AddUserResult(ctx context.Context, result UserResult)
	UpdateUserScore(ctx context.Context, score UserScore)
	AddWordleGroup(ctx context.Context, groupID string) // and users?
}
