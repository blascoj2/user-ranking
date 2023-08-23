package presenters

import (
	"github.com/user-ranking/internal/application"
	"github.com/user-ranking/internal/domain"
)

func ProvideJsonRankingPresenter() application.RankingPresenter {
	return &JsonRankingPresenter{}
}

type JsonRankingPresenter struct{}

type JsonUserRank struct {
	UserId string `json:"user_id"`
	Score  int    `json:"score"`
}

func (p *JsonRankingPresenter) PresentMany(ranking []domain.UserRank) interface{} {
	jsonRanking := make([]JsonUserRank, 0)
	for _, userRank := range ranking {
		jsonRanking = append(jsonRanking, JsonUserRank{
			UserId: string(userRank.UserId),
			Score:  int(userRank.Score),
		})
	}
	return jsonRanking
}
