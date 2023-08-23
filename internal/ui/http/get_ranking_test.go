package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/user-ranking/internal/application"
	"github.com/user-ranking/internal/infrastructure/persistence/inmemory"
	presenters "github.com/user-ranking/internal/ui/presenter"

	"github.com/stretchr/testify/assert"
)

func TestGetRankingAPI(t *testing.T) {
	t.Parallel()

	t.Run("should get top ranking successfully", func(t *testing.T) {
		controller := provideControllerTestSuite()

		createUsers(t, controller)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ranking?type=top100", nil)
		req.Header = map[string][]string{"Content-Type": {"application/x-www-form-urlencoded"}}
		controller.engine.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, `[{"user_id":"1","score":100},{"user_id":"3","score":90},{"user_id":"2","score":80},{"user_id":"4","score":70}]`, string(w.Body.Bytes()))
	})

	t.Run("should get relative ranking successfully (at2/1)", func(t *testing.T) {
		controller := provideControllerTestSuite()

		createUsers(t, controller)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ranking?type=at2/1", nil)
		req.Header = map[string][]string{"Content-Type": {"application/x-www-form-urlencoded"}}
		controller.engine.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, `[{"user_id":"1","score":100},{"user_id":"3","score":90},{"user_id":"2","score":80}]`, string(w.Body.Bytes()))
	})

	t.Run("should get relative ranking successfully (at4/1 last position)", func(t *testing.T) {
		controller := provideControllerTestSuite()

		createUsers(t, controller)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ranking?type=at4/1", nil)
		req.Header = map[string][]string{"Content-Type": {"application/x-www-form-urlencoded"}}
		controller.engine.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, `[{"user_id":"2","score":80},{"user_id":"4","score":70}]`, string(w.Body.Bytes()))
	})

	t.Run("should get relative ranking successfully (at1/1 first position)", func(t *testing.T) {
		controller := provideControllerTestSuite()

		createUsers(t, controller)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ranking?type=at1/1", nil)
		req.Header = map[string][]string{"Content-Type": {"application/x-www-form-urlencoded"}}
		controller.engine.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, `[{"user_id":"1","score":100},{"user_id":"3","score":90}]`, string(w.Body.Bytes()))
	})

	t.Run("should return invalid top ranking type", func(t *testing.T) {
		controller := provideControllerTestSuite()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ranking?type=top-100", nil)
		req.Header = map[string][]string{"Content-Type": {"application/x-www-form-urlencoded"}}
		controller.engine.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)
	})

	t.Run("should return invalid relative ranking type", func(t *testing.T) {
		controller := provideControllerTestSuite()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ranking?type=at10/", nil)
		req.Header = map[string][]string{"Content-Type": {"application/x-www-form-urlencoded"}}
		controller.engine.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)
	})
}

func createUsers(t *testing.T, controller *Controller) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/1/score", strings.NewReader(`{ "total":100}`))
	req.Header = map[string][]string{"Content-Type": {"application/json"}}
	controller.engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/user/2/score", strings.NewReader(`{ "total":80}`))
	req.Header = map[string][]string{"Content-Type": {"application/json"}}
	controller.engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/user/3/score", strings.NewReader(`{ "total":90}`))
	req.Header = map[string][]string{"Content-Type": {"application/json"}}
	controller.engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/user/4/score", strings.NewReader(`{ "total":70}`))
	req.Header = map[string][]string{"Content-Type": {"application/json"}}
	controller.engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func provideControllerTestSuite() *Controller {
	inmemoryRepository := inmemory.ProvideInMemoryRepository()
	presenter := presenters.ProvideJsonRankingPresenter()
	saveUserScoreHandler := application.ProvideSaveUserScoreHandler(inmemoryRepository)
	getRankingHandler := application.ProvideGetRankingHandler(inmemoryRepository, presenter)

	return ProvideController(saveUserScoreHandler, getRankingHandler)
}
