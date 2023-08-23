package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveUserScoreAPI(t *testing.T) {
	t.Parallel()

	t.Run("should create user score successfully", func(t *testing.T) {
		controller := provideControllerTestSuite()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/1/score", strings.NewReader(`{ "total":100}`))
		req.Header = map[string][]string{"Content-Type": {"application/json"}}
		controller.engine.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})

	t.Run("should update user partial score successfully", func(t *testing.T) {
		controller := provideControllerTestSuite()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/1/score", strings.NewReader(`{ "total":100}`))
		req.Header = map[string][]string{"Content-Type": {"application/json"}}
		controller.engine.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("POST", "/user/1/score", strings.NewReader(`{ "score":"+30"}`))
		req.Header = map[string][]string{"Content-Type": {"application/json"}}
		controller.engine.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/ranking?type=top10", nil)
		req.Header = map[string][]string{"Content-Type": {"application/x-www-form-urlencoded"}}
		controller.engine.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, `[{"user_id":"1","score":130}]`, string(w.Body.Bytes()))
	})

	t.Run("should return invalid partial score", func(t *testing.T) {
		controller := provideControllerTestSuite()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/1/score", strings.NewReader(`{ "score":"X100"}`))
		req.Header = map[string][]string{"Content-Type": {"application/json"}}
		controller.engine.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)
	})

	t.Run("should return user not found updating partial score", func(t *testing.T) {
		controller := provideControllerTestSuite()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/1/score", strings.NewReader(`{ "score":"+100"}`))
		req.Header = map[string][]string{"Content-Type": {"application/json"}}
		controller.engine.ServeHTTP(w, req)
		assert.Equal(t, 404, w.Code)
	})
}
