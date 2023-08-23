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

		userId := "1"

		id, err := domain.NewUserId(userId)
		assert.NoError(t, err)
		assert.Equal(t, userId, string(*id))
	})

	t.Run("invalid user id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userId := ""

		id, err := domain.NewUserId(userId)
		assert.Error(t, err)
		assert.Nil(t, id)
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

		id := domain.UserId("1")
		score := domain.Score(4)
		expectedRank := domain.UserRank{
			UserId: id,
			Score:  score,
		}

		rank := domain.NewUserRank(id, score)
		assert.Equal(t, expectedRank, rank)
	})
}
