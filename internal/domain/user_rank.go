package domain

import (
	"fmt"
	"strconv"
)

type UserId string
type Score int
type UserRank struct {
	UserId UserId
	Score  Score
}

func NewUserRank(userId UserId, score Score) UserRank {
	return UserRank{
		UserId: userId,
		Score:  score,
	}
}

func NewUserId(value string) (*UserId, error) {
	if value == "" {
		msg := fmt.Sprintf("invalid user id: %s", value)
		return nil, NewValidationErr(msg, InvalidRequestErrorTag)
	}

	uId := UserId(value)
	return &uId, nil
}

func NewScore(value int) Score {
	return Score(value)
}

func NewPartialScore(partialScore string) (*Score, error) {
	partialScoreInt, err := strconv.Atoi(partialScore)
	if err != nil {
		msg := fmt.Sprintf("invalid partial score format. Reason: %s", err)
		return nil, NewValidationErr(msg, InvalidRequestErrorTag)
	}

	score := Score(partialScoreInt)
	return &score, nil
}

func (u *UserRank) UpdateUserScore(partialScore Score) {
	u.Score += partialScore
}
