package application

import (
	"context"
	"github.com/user-ranking/internal/domain"
)

func ProvideSaveUserScoreHandler(
	userScoreRepository domain.Repository,
) SaveUserScoreHandler {
	return SaveUserScoreHandler{
		userScoreRepository: userScoreRepository,
	}
}

type SaveUserScoreHandler struct {
	userScoreRepository domain.Repository
}

type SaveUserScoreCommand struct {
	UserId       string
	PartialScore *string
	TotalScore   *int
}

func (h *SaveUserScoreHandler) Handle(ctx context.Context, cmd SaveUserScoreCommand) error {
	userId, err := domain.NewUserId(cmd.UserId)
	if err != nil {
		return err
	}

	if cmd.PartialScore != nil {
		partialScore, err := domain.NewPartialScore(*cmd.PartialScore)
		if err != nil {
			return err
		}

		return h.userScoreRepository.UpdateUserScore(ctx, *userId, *partialScore)
	}

	totalScore := domain.NewScore(*cmd.TotalScore)
	userRank := domain.NewUserRank(*userId, totalScore)
	return h.userScoreRepository.SaveUserScore(ctx, userRank)
}
