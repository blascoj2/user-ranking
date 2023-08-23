package inmemory_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/user-ranking/internal/domain"
	"github.com/user-ranking/internal/infrastructure/persistence/inmemory"
	"testing"
)

func TestUserRankSerializer_Serialize(t *testing.T) {
	t.Run("It should serialize user rank into an in memory document", func(t *testing.T) {
		serializer := inmemory.UserRankSerializer{}
		userRank := domain.UserRank{
			UserId: "1",
			Score:  100,
		}

		document := serializer.Serialize(userRank)
		assert.Equal(t, string(userRank.UserId), document.UserId)
		assert.Equal(t, int(userRank.Score), document.Score)
	})
}

func TestUserRankSerializer_Deserialize(t *testing.T) {
	t.Run("It should denormalize a user rank document into a user rank entity", func(t *testing.T) {
		serializer := inmemory.UserRankSerializer{}
		doc := inmemory.UserRankDocument{
			UserId: "1",
			Score:  100,
		}

		userRank := serializer.Deserialize(doc)
		assert.Equal(t, doc.UserId, string(userRank.UserId))
		assert.Equal(t, doc.Score, int(userRank.Score))
	})
}
