package application

import (
	"context"

	"github.com/user-ranking/internal/domain"
)

func ProvideGetRankingHandler(
	userScoreRepository domain.Repository,
	presenter RankingPresenter,
) GetRankingHandler {
	return GetRankingHandler{
		userScoreRepository: userScoreRepository,
		presenter:           presenter,
	}
}

type GetRankingHandler struct {
	userScoreRepository domain.Repository
	presenter           RankingPresenter
}

type GetRankingQuery struct {
	RankingType string
}

func (h *GetRankingHandler) Handle(ctx context.Context, query GetRankingQuery) (interface{}, error) {
	rankingType, err := domain.NewRankingType(query.RankingType)
	if err != nil {
		return nil, err
	}

	var ranking []domain.UserRank
	if rankingType.IsTop() {
		ranking, err = h.userScoreRepository.GetTopRanking(ctx, rankingType.Position)
		if err != nil {
			return nil, err
		}
	} else {
		ranking, err = h.userScoreRepository.GetRelativeRanking(ctx, rankingType.Position, rankingType.RangeSize)
		if err != nil {
			return nil, err
		}
	}

	return h.presenter.PresentMany(ranking), nil
}
