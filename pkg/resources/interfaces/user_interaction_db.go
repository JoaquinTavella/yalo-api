package interfaces

import "yalo-api/pkg/resources/models"

type UserInteractionDB interface {
	Set(key string, userInteraction models.UserInteraction) bool
	Get(key string) (models.UserInteraction, bool)
	GetAll() map[string]models.UserInteraction
	GetAllByUserID(userID string) map[string]models.UserInteraction
}
