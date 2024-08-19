package servicesmock

import (
	"errors"

	"yalo-api/internal/handlers/dtos"
)

type UserInteractionServiceMock struct {
	IsOk bool
}

//nolint:goerr113
func (s UserInteractionServiceMock) Save(_ []dtos.UserInteraction) error {
	if !s.IsOk {
		return errors.New("error saving user interaction")
	}

	return nil
}

//nolint:goerr113
func (s UserInteractionServiceMock) Get(_ string) (dtos.ProductRecommendationResponse, error) {
	if !s.IsOk {
		return dtos.ProductRecommendationResponse{}, errors.New("no interactions found")
	}

	return dtos.ProductRecommendationResponse{
		UserID:   "1",
		Products: []string{"1", "2", "3"},
	}, nil
}
