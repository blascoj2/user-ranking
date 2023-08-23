package application_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/user-ranking/internal/application"
	"github.com/user-ranking/internal/domain"
	application_mock "github.com/user-ranking/internal/mocks/application"
	domain_mock "github.com/user-ranking/internal/mocks/domain"
	"testing"
)

func TestGetRankingTestHandler_Handle(t *testing.T) {
	t.Parallel()

	t.Run("it should get top ranking successfully", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		suite := newGetRankingTestSuite(ctrl)

		expectedUserRank := domain.UserRank{
			UserId: domain.UserId("1"),
			Score:  domain.Score(100),
		}
		expectedRanking := []domain.UserRank{expectedUserRank}
		suite.repositoryMock.
			EXPECT().
			GetTopRanking(ctx, 100).
			Return(expectedRanking, nil).
			Times(1)

		expectedResponse := `{"foo": "bar"}`
		suite.presenterMock.
			EXPECT().
			PresentMany(expectedRanking).
			Return(expectedResponse).
			Times(1)

		command := application.GetRankingQuery{
			RankingType: "top100",
		}
		response, err := suite.handler.Handle(ctx, command)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("it should get relative ranking successfully", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		suite := newGetRankingTestSuite(ctrl)

		expectedUserRank := domain.UserRank{
			UserId: domain.UserId("1"),
			Score:  domain.Score(100),
		}
		expectedRanking := []domain.UserRank{expectedUserRank}
		suite.repositoryMock.
			EXPECT().
			GetRelativeRanking(ctx, 100, 3).
			Return(expectedRanking, nil).
			Times(1)

		expectedResponse := `{"foo": "bar"}`
		suite.presenterMock.
			EXPECT().
			PresentMany(expectedRanking).
			Return(expectedResponse).
			Times(1)

		command := application.GetRankingQuery{
			RankingType: "at100/3",
		}
		response, err := suite.handler.Handle(ctx, command)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("it should return invalid ranking type", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		suite := newGetRankingTestSuite(ctrl)

		command := application.GetRankingQuery{
			RankingType: "at100",
		}
		response, err := suite.handler.Handle(ctx, command)
		assert.Error(t, err)
		assert.IsType(t, err, domain.ValidationErr{})
		assert.Nil(t, response)
	})

	t.Run("it should return repository error getting top ranking", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		suite := newGetRankingTestSuite(ctrl)

		expectedErr := errors.New("testing error")
		suite.repositoryMock.
			EXPECT().
			GetTopRanking(ctx, gomock.Any()).
			Return(nil, expectedErr).
			Times(1)

		command := application.GetRankingQuery{
			RankingType: "top100",
		}
		response, err := suite.handler.Handle(ctx, command)
		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("it should return repository error getting relative ranking", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		suite := newGetRankingTestSuite(ctrl)

		expectedErr := errors.New("testing error")
		suite.repositoryMock.
			EXPECT().
			GetRelativeRanking(ctx, gomock.Any(), gomock.Any()).
			Return(nil, expectedErr).
			Times(1)

		command := application.GetRankingQuery{
			RankingType: "at100/3",
		}
		response, err := suite.handler.Handle(ctx, command)
		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

type GetRankingTestSuite struct {
	repositoryMock domain_mock.MockRepository
	presenterMock  application_mock.MockRankingPresenter
	handler        application.GetRankingHandler
}

func newGetRankingTestSuite(ctrl *gomock.Controller) GetRankingTestSuite {
	repositoryMock := domain_mock.NewMockRepository(ctrl)
	presenterMock := application_mock.NewMockRankingPresenter(ctrl)
	handler := application.ProvideGetRankingHandler(repositoryMock, presenterMock)

	return GetRankingTestSuite{
		repositoryMock: *repositoryMock,
		presenterMock:  *presenterMock,
		handler:        handler,
	}
}
