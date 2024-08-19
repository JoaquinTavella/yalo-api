package dtos

import (
	"time"
	"yalo-api/internal/utils"
)

type Actions string

const (
	View      Actions = "view"
	Click     Actions = "click"
	AddToCart Actions = "add_to_cart"
)

type UserInteraction struct {
	UserID               string    `json:"user_id"`
	ProductSKU           string    `json:"product_sku"`
	Action               Actions   `json:"action"`
	InteractionTimestamp time.Time `json:"interaction_timestamp"`
	InteractionDuration  int       `json:"interaction_duration,omitempty"`
}

type UserInteractionRequest struct {
	UserInteraction []UserInteraction `json:"user_interaction"`
}

type UserInteractionResponse struct {
	BaseResponse
	Message string `json:"message,omitempty"`
}

func (u UserInteractionRequest) Validate() error {
	if len(u.UserInteraction) == 0 {
		return utils.NewError(nil, "The body can't be empty", u.Validate)
	}

	for _, interaction := range u.UserInteraction {
		if interaction.UserID == "" || interaction.ProductSKU == "" || interaction.InteractionTimestamp.IsZero() {
			return utils.NewError(nil, "user_id, product_sku, and InteractionTimestamp are required", u.Validate)
		}

		if !isValidAction(interaction.Action) {
			return utils.NewError(nil, "Invalid action", u.Validate)
		}
	}

	return nil
}

func isValidAction(action Actions) bool {
	switch action {
	case View, Click, AddToCart:
		return true
	default:
		return false
	}
}
