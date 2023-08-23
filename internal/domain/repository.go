package domain

import "context"

type Repository interface {
	SaveUserScore(ctx context.Context, userScore UserRank) error
	UpdateUserScore(ctx context.Context, userId UserId, partialScore Score) error
	GetTopRanking(ctx context.Context, position int) ([]UserRank, error)
	GetRelativeRanking(ctx context.Context, position, rangeSize int) ([]UserRank, error)
}
