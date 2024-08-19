package resourcesmock

import (
	"time"

	"yalo-api/internal/handlers/dtos"
	"yalo-api/internal/resources/models"
)

type UserInteractionDBMock struct {
	ErrorOnSet       bool
	ErrorOnGet       bool
	ErrorOnGetByUser bool
	Return3          bool
}

func (u *UserInteractionDBMock) Set(_ string, _ models.UserInteraction) bool {
	return !(u.ErrorOnSet)
}

func (u *UserInteractionDBMock) Get(_ string) (models.UserInteraction, bool) {
	if u.ErrorOnGet {
		return models.UserInteraction{}, false
	}

	return models.UserInteraction{}, true
}

func (u *UserInteractionDBMock) Delete(_ string) {
}

func (u *UserInteractionDBMock) GetAll() map[string]models.UserInteraction {
	return make(map[string]models.UserInteraction)
}

//nolint:gomnd
func (u *UserInteractionDBMock) GetAllByUserID(_ string) map[string]models.UserInteraction {
	if u.ErrorOnGetByUser {
		return nil
	}

	if u.Return3 {
		savedMap := map[string]models.UserInteraction{
			"1": {
				UserID:               "1",
				ProductSKU:           "1",
				Action:               dtos.AddToCart,
				InteractionTimestamp: time.Now(),
				InteractionDuration:  1,
			},
			"2": {
				UserID:               "1",
				ProductSKU:           "2",
				Action:               dtos.AddToCart,
				InteractionTimestamp: time.Now(),
				InteractionDuration:  2,
			},
			"3": {
				UserID:               "1",
				ProductSKU:           "3",
				Action:               dtos.AddToCart,
				InteractionTimestamp: time.Now(),
				InteractionDuration:  4,
			},
		}

		return savedMap
	}

	return map[string]models.UserInteraction{
		"1": {
			UserID:               "1",
			ProductSKU:           "1",
			Action:               dtos.AddToCart,
			InteractionTimestamp: time.Now(),
			InteractionDuration:  1,
		},
	}
}
