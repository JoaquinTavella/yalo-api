package services

import (
	"fmt"
	"sort"
	"time"

	"yalo-api/internal/handlers/dtos"
	"yalo-api/internal/resources/interfaces"
	"yalo-api/internal/resources/models"
	"yalo-api/internal/utils"
)

type UserInteractionService struct {
	DataInteractionRepository interfaces.UserInteractionDB
}

const MinAmountInteractions = 3

func NewUserInteractionService(dataInteractionRepository interfaces.UserInteractionDB) UserInteractionService {
	return UserInteractionService{
		DataInteractionRepository: dataInteractionRepository,
	}
}

func (s UserInteractionService) Save(userInteraction []dtos.UserInteraction) error {
	for _, interaction := range userInteraction {
		userInteractionModel := models.UserInteraction{
			UserID:               interaction.UserID,
			ProductSKU:           interaction.ProductSKU,
			Action:               interaction.Action,
			InteractionTimestamp: interaction.InteractionTimestamp,
			InteractionDuration:  interaction.InteractionDuration,
		}

		key := CreateKey(interaction.UserID, interaction.InteractionTimestamp, interaction.Action, interaction.ProductSKU)
		set := s.DataInteractionRepository.Set(key, userInteractionModel)

		if !set {
			return utils.NewError(nil, "could not save", s)
		}
	}

	return nil
}

func CreateKey(userID string, interactionTimestamp time.Time, action dtos.Actions, product string) string {
	return fmt.Sprintf("%s-%s-%s-%s", userID, interactionTimestamp, action, product)
}

func (s UserInteractionService) Get(userID string) (dtos.ProductRecommendationResponse, error) {
	userInteraction := s.DataInteractionRepository.GetAllByUserID(userID)
	if len(userInteraction) == 0 {
		return dtos.ProductRecommendationResponse{
			UserID:   userID,
			Products: []string{},
		}, utils.NewError(nil, "no interactions found", s)
	}

	return dtos.ProductRecommendationResponse{
		UserID:   userID,
		Products: calculateScore(userInteraction),
	}, nil
}

func calculateScore(userInteraction map[string]models.UserInteraction) []string {
	productScores := make(map[string]int)
	for _, interaction := range userInteraction {
		productScores[interaction.ProductSKU]++
	}

	return orderProductsByScore(productScores)
}

func orderProductsByScore(productScores map[string]int) []string {
	keys := make([]string, 0, len(productScores))
	for key := range productScores {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return productScores[keys[i]] > productScores[keys[j]]
	})

	if len(keys) < MinAmountInteractions {
		return keys
	}

	return keys[:MinAmountInteractions]
}
