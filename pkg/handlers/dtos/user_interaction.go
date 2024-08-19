package dtos

import "time"

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
