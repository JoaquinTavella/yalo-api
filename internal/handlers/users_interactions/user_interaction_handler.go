package handlers

import (
	"net/http"
	"yalo-api/internal/middlewares"

	"github.com/gin-gonic/gin"
	"yalo-api/internal/handlers/dtos"
	"yalo-api/internal/services/interfaces"
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
	logger := middlewares.LoggerFromContext(ginContext)

	var userInteractions dtos.UserInteractionRequest

	if err := ginContext.BindJSON(&userInteractions); err != nil {
		responseError := dtos.UserInteractionResponse{
			Message:      "Error parsing request",
			BaseResponse: dtos.BaseResponse{Errors: []interface{}{err.Error()}},
		}
		logger.Error().Err(err).Msgf("%s - Error: %s", responseError.Message, err.Error())
		ginContext.JSON(http.StatusBadRequest, responseError)

		return
	}

	err := userInteractions.Validate()
	if err != nil {
		responseError := dtos.UserInteractionResponse{
			Message:      "Invalid input",
			BaseResponse: dtos.BaseResponse{Errors: []interface{}{err.Error()}},
		}

		logger.Error().Err(err).Msgf("%s - Error: %s", responseError.Message, err.Error())

		ginContext.JSON(http.StatusBadRequest, responseError)

		return
	}

	err = h.UserInteractionService.Save(userInteractions.UserInteraction)
	if err != nil {
		responseError := dtos.UserInteractionResponse{
			Message:      "Error saving user interactions",
			BaseResponse: dtos.BaseResponse{Errors: []interface{}{err.Error()}},
		}

		logger.Error().Err(err).Msgf("%s - Error: %s", responseError.Message, err.Error())

		ginContext.JSON(http.StatusInternalServerError, responseError)

		return
	}

	ginContext.JSON(http.StatusOK, dtos.UserInteractionResponse{
		Message: "User interactions saved successfully",
	})
}

func (h UserInteractionHandler) GetUserInteraction(ginContext *gin.Context) {
	logger := middlewares.LoggerFromContext(ginContext)

	userID := ginContext.Param("user_id")

	productRecommendation, err := h.UserInteractionService.Get(userID)
	if err != nil {
		responseError := dtos.ProductRecommendationResponse{
			BaseResponse: dtos.BaseResponse{Errors: []interface{}{err.Error()}},
		}

		logger.Error().Err(err).Msgf("Error getting user recommendation: %s", err.Error())

		ginContext.JSON(http.StatusNotFound, responseError)

		return
	}

	ginContext.JSON(http.StatusOK, productRecommendation)
}
