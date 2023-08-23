package http

import (
	"github.com/gin-gonic/gin"
	"github.com/user-ranking/internal/application"
	"net/http"
)

type GetRankingHttpQuery struct {
	RankingType string `form:"type"`
}

func (c *Controller) getRanking(ctx *gin.Context) {
	var httpQuery GetRankingHttpQuery
	if err := ctx.ShouldBindQuery(&httpQuery); err != nil {
		c.logger.Errorf("binding get ranking query, reason: %s", err)
		c.setBadRequestResponse(ctx, err)
		return
	}

	ranking, err := c.getRankingHandler.Handle(ctx, application.GetRankingQuery{
		RankingType: httpQuery.RankingType,
	})
	if err != nil {
		c.logger.Errorf("getting ranking, reason: %s", err)
		c.setErrorResponse(ctx, err)
		return
	}

	c.setSuccessResponse(ctx, http.StatusOK, ranking)
}
