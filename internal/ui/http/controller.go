package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/user-ranking/internal/application"
	"github.com/user-ranking/internal/domain"
)

func ProvideController(
	saveUserScoreHandler application.SaveUserScoreHandler,
	getRankingHandler application.GetRankingHandler,
) *Controller {
	logger := log.WithFields(log.Fields{"file": "controller", "service": "user-ranking"})

	engine := gin.Default()
	engine.HandleMethodNotAllowed = true
	c := &Controller{
		logger:               logger,
		engine:               engine,
		saveUserScoreHandler: saveUserScoreHandler,
		getRankingHandler:    getRankingHandler,
	}

	c.Setup()
	return c
}

type Controller struct {
	logger               *log.Entry
	engine               *gin.Engine
	saveUserScoreHandler application.SaveUserScoreHandler
	getRankingHandler    application.GetRankingHandler
}

func (c *Controller) Setup() {
	c.engine.POST("/user/:user_id/score", c.saveUserScore)
	c.engine.GET("/ranking", c.getRanking)
}

func (c *Controller) Run() {
	c.engine.Run("0.0.0.0:8080")
}

func (c *Controller) setSuccessResponse(ctx *gin.Context, statusCode int, body interface{}) {
	if body != nil {
		ctx.JSON(statusCode, body)
		return
	}

	ctx.Status(statusCode)
}

func (c *Controller) setBadRequestResponse(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, NewHttpError(err.Error(), domain.BadRequestCode, domain.InvalidRequestErrorTag))
}

func (c *Controller) setInternalErrorResponse(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, NewHttpError(err.Error(), domain.UnexpectedErrorCode, domain.ServerErrorTag))
}

func (c *Controller) setErrorResponse(ctx *gin.Context, err error) {
	if domainErr, ok := err.(domain.ErrInterface); ok {
		statusCode := http.StatusInternalServerError
		if errors.As(err, &domain.ValidationErr{}) {
			statusCode = http.StatusBadRequest
		}

		if errors.As(err, &domain.NotFoundErr{}) {
			statusCode = http.StatusNotFound
		}

		ctx.JSON(statusCode, NewHttpError(domainErr.Message(), domainErr.Code(), domainErr.Tag()))
		return
	}

	c.setInternalErrorResponse(ctx, err)
	return
}

func (c *Controller) setInvalidContentType(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusUnsupportedMediaType)
}
