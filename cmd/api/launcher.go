package main

import (
	"github.com/gin-gonic/gin"
	"github.com/user-ranking/internal/infrastructure/flags"
	logging "github.com/user-ranking/internal/infrastructure/logger"
	"github.com/user-ranking/internal/ui/http"
)

func ProvideLauncher(
	flags flags.Flags,
	httpController *http.Controller,
) Launcher {
	return Launcher{
		flags:          flags,
		httpController: httpController,
	}
}

type Launcher struct {
	flags          flags.Flags
	httpController *http.Controller
}

func (l *Launcher) Launch() {
	gin.SetMode(gin.ReleaseMode)
	logging.InitLogger(logging.GetTraceLevelWithName(l.flags.LoggingFlags.TraceLevel))

	l.httpController.Run()
}
