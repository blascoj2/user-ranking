package application

import "github.com/user-ranking/internal/domain"

type RankingPresenter interface {
	PresentMany(ranking []domain.UserRank) interface{}
}
