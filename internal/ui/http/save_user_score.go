package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/user-ranking/internal/application"
	"net/http"
)

type SaveUserScoreHttpRequest struct {
	PartialScore *string `json:"score"`
	TotalScore   *int    `json:"total"`
}

func (c *Controller) saveUserScore(ctx *gin.Context) {
	if ctx.ContentType() != "application/json" {
		c.logger.Errorf("save user score: unsupported media type")
		ctx.AbortWithStatus(http.StatusUnsupportedMediaType)
		return
	}

	userId := ctx.Param("user_id")
	var httpRequest SaveUserScoreHttpRequest
	if err := ctx.ShouldBindJSON(&httpRequest); err != nil {
		c.logger.Errorf("binding save user score request, reason: %s", err)
		c.setBadRequestResponse(ctx, err)
		return
	}

	if httpRequest.PartialScore == nil && httpRequest.TotalScore == nil {
		msg := "save user request validation error, reason: one of partial or total score must be informed"
		c.logger.Errorf(msg)
		c.setBadRequestResponse(ctx, errors.New(msg))
		return
	}

	if httpRequest.PartialScore != nil && httpRequest.TotalScore != nil {
		msg := "save user request validation error, reason: only one of partial or total score can be informed"
		c.logger.Errorf(msg)
		c.setBadRequestResponse(ctx, errors.New(msg))
		return
	}

	if err := c.saveUserScoreHandler.Handle(ctx, application.SaveUserScoreCommand{
		UserId:       userId,
		PartialScore: httpRequest.PartialScore,
		TotalScore:   httpRequest.TotalScore,
	}); err != nil {
		c.logger.Errorf("saving user score, reason: %s", err)
		c.setErrorResponse(ctx, err)
		return
	}

	c.setSuccessResponse(ctx, http.StatusOK, nil)
}
