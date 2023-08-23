package presenters_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/user-ranking/internal/domain"
	presenters "github.com/user-ranking/internal/ui/presenter"
)

func TestJsonRankingPresenter_PresentMany(t *testing.T) {

	userRank := domain.UserRank{
		UserId: "1",
		Score:  100,
	}

	ranking := []domain.UserRank{userRank}
	p := presenters.JsonRankingPresenter{}
	dto := p.PresentMany(ranking)

	jsonBytes, err := json.Marshal(dto)
	assert.NoError(t, err)

	expectedJson := []byte(`[{
	  	"user_id": "1",
	  	"score": 100
	}]`)

	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, expectedJson)
	assert.NoError(t, err)

	assert.Equal(t, buffer.String(), string(jsonBytes))
}
