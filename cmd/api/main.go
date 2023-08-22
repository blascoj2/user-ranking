package main

import (
	flags2 "github.com/user-ranking/internal/infrastructure/flags"
	"github.com/user-ranking/internal/ui/http"
)

func main() {
	flags := flags2.ProvideFlags()
	httpController := http.ProvideController()
	launcher := ProvideLauncher(flags, httpController)
	launcher.Launch()
}
