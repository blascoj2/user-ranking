package domain

import (
	"fmt"
	"strconv"
	"strings"
)

type Type int

const (
	Top      Type = 0
	Relative Type = 1
)

type RankingType struct {
	Type      Type
	Position  int
	RangeSize int
}

func NewRankingType(rankingType string) (*RankingType, error) {
	rankingType = strings.ToLower(rankingType)
	rankingVO := &RankingType{}

	if strings.HasPrefix(rankingType, "top") {
		rankingVO.Type = Top
	} else if strings.HasPrefix(rankingType, "at") {
		rankingVO.Type = Relative
	} else {
		msg := fmt.Sprintf("invalid ranking type. Allowed types: topX, atX/X.")
		return nil, NewValidationErr(msg, InvalidRequestErrorTag)
	}

	if rankingVO.IsTop() {
		positionStr := strings.TrimPrefix(rankingType, "top")

		position, err := strconv.Atoi(positionStr)
		if err != nil {
			msg := fmt.Sprintf("invalid ranking type. Invalid top position format: %s.", positionStr)
			return nil, NewValidationErr(msg, InvalidRequestErrorTag)
		}

		rankingVO.Position = position
		return rankingVO.assert()
	}

	positionRange := strings.SplitN(strings.TrimPrefix(rankingType, "at"), "/", 2)
	if len(positionRange) != 2 {
		msg := fmt.Sprintf("invalid ranking type. Relative position range must conatin position and rangeSize: %s.", positionRange)
		return nil, NewValidationErr(msg, InvalidRequestErrorTag)
	}

	position, err := strconv.Atoi(positionRange[0])
	if err != nil {
		msg := fmt.Sprintf("invalid ranking type. Invalid relative position format: %s.", positionRange[0])
		return nil, NewValidationErr(msg, InvalidRequestErrorTag)
	}

	rangeSize, err := strconv.Atoi(positionRange[1])
	if err != nil {
		msg := fmt.Sprintf("invalid ranking type. Invalid relative range size format: %s.", positionRange[1])
		return nil, NewValidationErr(msg, InvalidRequestErrorTag)
	}

	rankingVO.Position = position
	rankingVO.RangeSize = rangeSize
	return rankingVO.assert()
}

func (r *RankingType) assert() (*RankingType, error) {
	if r.Position <= 0 {
		msg := fmt.Sprintf("invalid ranking type position. Must be greater than 0")
		return nil, NewValidationErr(msg, InvalidRequestErrorTag)
	}

	if !r.IsTop() && r.RangeSize < 0 {
		msg := fmt.Sprintf("invalid relative ranking type range size. Must be greater equals than 0")
		return nil, NewValidationErr(msg, InvalidRequestErrorTag)
	}

	return r, nil
}

func (r *RankingType) IsTop() bool {
	return r.Type == Top
}
