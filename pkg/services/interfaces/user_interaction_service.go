package interfaces

import (
	"yalo-api/pkg/handlers/dtos"
)

type SaveUserInteractionI interface {
	Save(_ []dtos.UserInteraction) error
}

type GetUserInteractionI interface {
	Get(_ string) (dtos.ProductRecommendationResponse, error)
}

type UserInteractionServiceInterface interface {
	SaveUserInteractionI
	GetUserInteractionI
}
