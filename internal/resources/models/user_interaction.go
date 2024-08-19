package models

import (
	"time"

	"yalo-api/internal/handlers/dtos"
)

type UserInteraction struct {
	UserID               string       `json:"user_id"`
	ProductSKU           string       `json:"product_sku"`
	Action               dtos.Actions `json:"action"`
	InteractionTimestamp time.Time    `json:"interaction_timestamp"`
	InteractionDuration  int          `json:"interaction_duration,omitempty"`
}
