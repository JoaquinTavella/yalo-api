package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yalo-api/pkg/handlers/dtos"
	"yalo-api/pkg/services/interfaces"
)

type UserInteractionHandler struct {
	UserInteractionService interfaces.UserInteractionServiceInterface
}

func NewUserInteractionHandler(userInteractionService interfaces.UserInteractionServiceInterface) UserInteractionHandler {
	return UserInteractionHandler{
		UserInteractionService: userInteractionService,
	}
}

func (h UserInteractionHandler) SaveUserInteraction(ginContext *gin.Context) {
	var userInteractions []dtos.UserInteraction

	if err := ginContext.BindJSON(&userInteractions); err != nil {
		responseError := dtos.UserInteractionResponse{
			Message:      "Error parsing request",
			BaseResponse: dtos.BaseResponse{Errors: []interface{}{err.Error()}},
		}

		ginContext.JSON(http.StatusBadRequest, responseError)

		return
	}

	err := h.UserInteractionService.Save(userInteractions)
	if err != nil {
		responseError := dtos.UserInteractionResponse{
			Message:      "Error saving user interactions",
			BaseResponse: dtos.BaseResponse{Errors: []interface{}{err.Error()}},
		}

		ginContext.JSON(http.StatusInternalServerError, responseError)

		return
	}

	ginContext.JSON(http.StatusOK, dtos.UserInteractionResponse{
		Message: "User interactions saved successfully",
	})
}

func (h UserInteractionHandler) GetUserInteraction(ginContext *gin.Context) {
	userID := ginContext.Param("user_id")

	productRecommendation, err := h.UserInteractionService.Get(userID)
	if err != nil {
		responseError := dtos.ProductRecommendationResponse{
			BaseResponse: dtos.BaseResponse{Errors: []interface{}{err.Error()}},
		}

		ginContext.JSON(http.StatusNotFound, responseError)

		return
	}

	ginContext.JSON(http.StatusOK, productRecommendation)
}
