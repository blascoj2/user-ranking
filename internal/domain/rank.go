package domain

type UserId uint
type Score int
type Rank struct {
	UserId UserId
	Score  Score
}

func NewRank(rawUserId uint, rawScore int) Rank {
	id := NewUserId(rawUserId)
	score := NewScore(rawScore)

	return Rank{
		UserId: id,
		Score:  score,
	}
}

func NewUserId(value uint) UserId {
	return UserId(value)
}

func NewScore(value int) Score {
	return Score(value)
}
