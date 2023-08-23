package inmemory

import "github.com/user-ranking/internal/domain"

type UserRankSerializer struct{}

type UserRankDocument struct {
	UserId string
	Score  int
}

func (s UserRankSerializer) Serialize(userRank domain.UserRank) UserRankDocument {
	return UserRankDocument{
		UserId: string(userRank.UserId),
		Score:  int(userRank.Score),
	}
}

func (s UserRankSerializer) Deserialize(document UserRankDocument) domain.UserRank {
	return domain.UserRank{
		UserId: domain.UserId(document.UserId),
		Score:  domain.Score(document.Score),
	}
}
