package http

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ProvideController() *Controller {
	logger := log.WithFields(log.Fields{"file": "controller", "service": "car-pooling"})

	engine := gin.Default()
	engine.HandleMethodNotAllowed = true
	c := &Controller{
		logger: logger,
		engine: engine,
	}

	c.Setup()
	return c
}

type Controller struct {
	logger *log.Entry
	engine *gin.Engine
}

func (c *Controller) Setup() {
	//c.engine.GET("/status", c.status)
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

func (c *Controller) setInvalidContentType(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusUnsupportedMediaType)
}
