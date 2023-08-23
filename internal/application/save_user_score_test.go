package application_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/user-ranking/internal/application"
	"github.com/user-ranking/internal/domain"
	domain_mock "github.com/user-ranking/internal/mocks/domain"
	"strconv"
	"testing"
)

func TestSaveUserScoreTestHandler_Handle(t *testing.T) {
	t.Parallel()

	t.Run("it should save user total score successfully", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		suite := newSaveUserScoreTestSuite(ctrl)

		expectedUserRank := domain.UserRank{
			UserId: domain.UserId("1"),
			Score:  domain.Score(100),
		}
		suite.repositoryMock.
			EXPECT().
			SaveUserScore(ctx, expectedUserRank).
			Return(nil).
			Times(1)

		totalScore := int(expectedUserRank.Score)
		command := application.SaveUserScoreCommand{
			UserId:     string(expectedUserRank.UserId),
			TotalScore: &totalScore,
		}
		err := suite.handler.Handle(ctx, command)
		assert.NoError(t, err)
	})

	t.Run("it should save user partial score successfully", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		suite := newSaveUserScoreTestSuite(ctrl)

		userId := domain.UserId("1")
		score := domain.Score(-3)
		suite.repositoryMock.
			EXPECT().
			UpdateUserScore(ctx, userId, score).
			Return(nil).
			Times(1)

		partialScore := strconv.Itoa(int(score))
		command := application.SaveUserScoreCommand{
			UserId:       string(userId),
			PartialScore: &partialScore,
		}
		err := suite.handler.Handle(ctx, command)
		assert.NoError(t, err)
	})

	t.Run("it should return invalid user id", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		suite := newSaveUserScoreTestSuite(ctrl)

		partialScore := "+100"
		command := application.SaveUserScoreCommand{
			UserId:       "",
			PartialScore: &partialScore,
		}
		err := suite.handler.Handle(ctx, command)
		assert.Error(t, err)
		assert.IsType(t, err, domain.ValidationErr{})
	})

	t.Run("it should return invalid partial score", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		suite := newSaveUserScoreTestSuite(ctrl)

		partialScore := "x100"
		command := application.SaveUserScoreCommand{
			UserId:       "1",
			PartialScore: &partialScore,
		}
		err := suite.handler.Handle(ctx, command)
		assert.Error(t, err)
		assert.IsType(t, err, domain.ValidationErr{})
	})

	t.Run("it should error saving user total score", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		suite := newSaveUserScoreTestSuite(ctrl)

		expectedUserRank := domain.UserRank{
			UserId: domain.UserId("1"),
			Score:  domain.Score(100),
		}
		suite.repositoryMock.
			EXPECT().
			SaveUserScore(ctx, gomock.Any()).
			Return(errors.New("testing error")).
			Times(1)

		totalScore := int(expectedUserRank.Score)
		command := application.SaveUserScoreCommand{
			UserId:     string(expectedUserRank.UserId),
			TotalScore: &totalScore,
		}
		err := suite.handler.Handle(ctx, command)
		assert.Error(t, err)
	})

	t.Run("it should return error saving user partial score", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		suite := newSaveUserScoreTestSuite(ctrl)

		userId := domain.UserId("1")
		score := domain.Score(-3)
		suite.repositoryMock.
			EXPECT().
			UpdateUserScore(ctx, gomock.Any(), gomock.Any()).
			Return(errors.New("testing error")).
			Times(1)

		partialScore := strconv.Itoa(int(score))
		command := application.SaveUserScoreCommand{
			UserId:       string(userId),
			PartialScore: &partialScore,
		}
		err := suite.handler.Handle(ctx, command)
		assert.Error(t, err)
	})
}

type SaveUserScoreTestSuite struct {
	repositoryMock domain_mock.MockRepository
	handler        application.SaveUserScoreHandler
}

func newSaveUserScoreTestSuite(ctrl *gomock.Controller) SaveUserScoreTestSuite {
	repositoryMock := domain_mock.NewMockRepository(ctrl)
	handler := application.ProvideSaveUserScoreHandler(repositoryMock)

	return SaveUserScoreTestSuite{
		repositoryMock: *repositoryMock,
		handler:        handler,
	}
}
