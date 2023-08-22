package domain_test

import (
	"github.com/user-ranking/internal/domain"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewUserId(t *testing.T) {
	t.Parallel()

	t.Run("create user id successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userId := uint(0)

		id := domain.NewUserId(userId)
		assert.Equal(t, userId, uint(id))
	})
}

func TestNewScore(t *testing.T) {
	t.Parallel()

	t.Run("create score successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rawScore := 4

		score := domain.NewScore(rawScore)
		assert.Equal(t, rawScore, int(score))
	})
}

func TestNewRank(t *testing.T) {
	t.Parallel()

	t.Run("create rank successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := uint(0)
		score := int(4)
		expectedRank := domain.Rank{
			UserId: domain.UserId(id),
			Score:  domain.Score(score),
		}

		rank := domain.NewRank(id, score)
		assert.Equal(t, expectedRank, rank)
	})
}
