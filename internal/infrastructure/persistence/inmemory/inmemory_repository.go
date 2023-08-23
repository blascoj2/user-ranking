package inmemory

import (
	"container/heap"
	"context"
	"github.com/user-ranking/internal/domain"
	"sort"
	"sync"
)

func ProvideInMemoryRepository() domain.Repository {
	return &InMemoryRepository{
		usersRank:  make(map[string]UserRankDocument, 0),
		serializer: UserRankSerializer{},
	}
}

type RankingMaxHeap []UserRankDocument

func (h RankingMaxHeap) Len() int            { return len(h) }
func (h RankingMaxHeap) Less(i, j int) bool  { return h[i].Score > h[j].Score }
func (h RankingMaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *RankingMaxHeap) Push(x interface{}) { *h = append(*h, x.(UserRankDocument)) }
func (h *RankingMaxHeap) Pop() interface{} {
	oldHeap := *h
	n := len(oldHeap)
	x := oldHeap[n-1]
	*h = oldHeap[:n-1]
	return x
}
func (h *RankingMaxHeap) Insert(userRank UserRankDocument) {
	if len(*h) == 0 {
		*h = append(*h, userRank)
	} else {
		index := sort.Search(len(*h), func(i int) bool {
			return userRank.Score > (*h)[i].Score
		})
		*h = append(*h, UserRankDocument{})
		copy((*h)[index+1:], (*h)[index:])
		(*h)[index] = userRank
	}
}
func (h *RankingMaxHeap) UpdateScore(position int, userRank UserRankDocument) {
	heap.Remove(h, position)
	h.Insert(userRank)
}

type InMemoryRepository struct {
	mu         sync.RWMutex
	usersRank  map[string]UserRankDocument
	ranking    RankingMaxHeap
	serializer UserRankSerializer
}

func (r *InMemoryRepository) SaveUserScore(ctx context.Context, userRank domain.UserRank) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.saveUserRank(ctx, userRank)
}

func (r *InMemoryRepository) UpdateUserScore(ctx context.Context, userId domain.UserId, partialScore domain.Score) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	userRank, err := r.findUserRank(ctx, userId)
	if err != nil {
		return err
	}

	userRank.UpdateUserScore(partialScore)

	return r.saveUserRank(ctx, *userRank)
}

func (r *InMemoryRepository) GetTopRanking(_ context.Context, position int) ([]domain.UserRank, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if position > len(r.ranking) {
		position = len(r.ranking)
	}

	rankingDoc := r.ranking[:position]
	ranking := make([]domain.UserRank, 0)
	for _, userRankDoc := range rankingDoc {
		ranking = append(ranking, r.serializer.Deserialize(userRankDoc))
	}
	return ranking, nil
}

func (r *InMemoryRepository) GetRelativeRanking(_ context.Context, position, rangeSize int) ([]domain.UserRank, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if position > len(r.ranking) {
		position = len(r.ranking)
	}

	relativePosition := position - 1
	if relativePosition < 0 {
		relativePosition = 0
	}

	startIndex := relativePosition - rangeSize
	endIndex := (relativePosition + rangeSize) + 1

	if startIndex < 0 {
		startIndex = 0
	}

	if endIndex > len(r.ranking) {
		endIndex = len(r.ranking)
	}

	rankingDoc := r.ranking[startIndex:endIndex]
	ranking := make([]domain.UserRank, 0)
	for _, userRankDoc := range rankingDoc {
		ranking = append(ranking, r.serializer.Deserialize(userRankDoc))
	}
	return ranking, nil
}

func (r *InMemoryRepository) findUserRank(_ context.Context, userId domain.UserId) (*domain.UserRank, error) {
	userRankDoc, ok := r.usersRank[string(userId)]
	if !ok {
		return nil, domain.NewUserNotFoundErr(userId)
	}

	userRank := r.serializer.Deserialize(userRankDoc)
	return &userRank, nil
}

func (r *InMemoryRepository) saveUserRank(ctx context.Context, userRank domain.UserRank) error {
	userRankDoc := r.serializer.Serialize(userRank)
	r.usersRank[string(userRank.UserId)] = userRankDoc
	currentRankingPosition := r.findRankingIndex(ctx, userRank.UserId)

	if currentRankingPosition >= 0 {
		r.ranking.UpdateScore(currentRankingPosition, userRankDoc)
	} else {
		r.ranking.Insert(userRankDoc)
	}

	return nil
}

func (r *InMemoryRepository) findRankingIndex(_ context.Context, userId domain.UserId) int {
	for i, user := range r.ranking {
		if user.UserId == string(userId) {
			return i
		}
	}

	return -1
}
