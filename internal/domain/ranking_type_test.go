package domain_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/user-ranking/internal/domain"
	"testing"
)

func TestNewRankingType(t *testing.T) {
	t.Parallel()

	t.Run("create top ranking type successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rawRankingType := "top100"

		rankingType, err := domain.NewRankingType(rawRankingType)
		assert.NoError(t, err)
		assert.True(t, rankingType.IsTop())
		assert.Equal(t, rankingType.Position, 100)
	})

	t.Run("create relative ranking type successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rawRankingType := "at100/3"

		rankingType, err := domain.NewRankingType(rawRankingType)
		assert.NoError(t, err)
		assert.False(t, rankingType.IsTop())
		assert.Equal(t, rankingType.Position, 100)
		assert.Equal(t, rankingType.RangeSize, 3)
	})

	t.Run("invalid ranking type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rawRankingType := "invalid-ranking-type"

		rankingType, err := domain.NewRankingType(rawRankingType)
		assert.Error(t, err)
		assert.IsType(t, err, domain.ValidationErr{})
		assert.Nil(t, rankingType)
	})

	t.Run("invalid top ranking type position (negative position)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rawRankingType := "top-100"

		rankingType, err := domain.NewRankingType(rawRankingType)
		assert.Error(t, err)
		assert.IsType(t, err, domain.ValidationErr{})
		assert.Nil(t, rankingType)
	})

	t.Run("invalid top ranking type position(non numeric position)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rawRankingType := "topX"

		rankingType, err := domain.NewRankingType(rawRankingType)
		assert.Error(t, err)
		assert.IsType(t, err, domain.ValidationErr{})
		assert.Nil(t, rankingType)
	})

	t.Run("invalid relative ranking type position (negative position)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rawRankingType := "at-100"

		rankingType, err := domain.NewRankingType(rawRankingType)
		assert.Error(t, err)
		assert.IsType(t, err, domain.ValidationErr{})
		assert.Nil(t, rankingType)
	})

	t.Run("invalid relative ranking type position(non numeric position)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rawRankingType := "atX/X"

		rankingType, err := domain.NewRankingType(rawRankingType)
		assert.Error(t, err)
		assert.IsType(t, err, domain.ValidationErr{})
		assert.Nil(t, rankingType)
	})

	t.Run("invalid relative ranking type range size(no range size)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rawRankingType := "atX/"

		rankingType, err := domain.NewRankingType(rawRankingType)
		assert.Error(t, err)
		assert.IsType(t, err, domain.ValidationErr{})
		assert.Nil(t, rankingType)
	})

	t.Run("invalid relative ranking type range size (negative range size)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rawRankingType := "at100/-10"

		rankingType, err := domain.NewRankingType(rawRankingType)
		assert.Error(t, err)
		assert.IsType(t, err, domain.ValidationErr{})
		assert.Nil(t, rankingType)
	})

	t.Run("invalid top ranking type range size(non numeric range size)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rawRankingType := "at100/X"

		rankingType, err := domain.NewRankingType(rawRankingType)
		assert.Error(t, err)
		assert.IsType(t, err, domain.ValidationErr{})
		assert.Nil(t, rankingType)
	})
}
