package main

import (
	"github.com/user-ranking/internal/application"
	flags2 "github.com/user-ranking/internal/infrastructure/flags"
	"github.com/user-ranking/internal/infrastructure/persistence/inmemory"
	"github.com/user-ranking/internal/ui/http"
	presenters "github.com/user-ranking/internal/ui/presenter"
)

func main() {
	flags := flags2.ProvideFlags()
	inmemoryRepository := inmemory.ProvideInMemoryRepository()
	presenter := presenters.ProvideJsonRankingPresenter()
	saveUserScoreHandler := application.ProvideSaveUserScoreHandler(inmemoryRepository)
	getRankingHandler := application.ProvideGetRankingHandler(inmemoryRepository, presenter)
	httpController := http.ProvideController(saveUserScoreHandler, getRankingHandler)
	launcher := ProvideLauncher(flags, httpController)
	launcher.Launch()
}
