package dtos

type ProductRecommendationResponse struct {
	BaseResponse
	UserID   string   `json:"user_id"`
	Products []string `json:"products"`
}
