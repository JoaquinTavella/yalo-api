package handlers_test

//nolint:goimports
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	services_mock "yalo-api/mocks/services"
	"yalo-api/pkg/handlers/dtos"
	handlers "yalo-api/pkg/handlers/users_interactions"
)

//nolint:funlen
func TestSaveUserInteraction(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Parallel()

	t.Run("Successful Save", func(t *testing.T) {
		t.Parallel()

		writer := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(writer)
		handler := handlers.NewUserInteractionHandler(services_mock.UserInteractionServiceMock{IsOk: true})

		// Sample input
		input := []dtos.UserInteraction{
			{
				UserID:               "1",
				Action:               dtos.AddToCart,
				ProductSKU:           "123",
				InteractionTimestamp: time.Now(),
				InteractionDuration:  10,
			},
		}
		jsonValue, _ := json.Marshal(input)

		context.Request = httptest.NewRequest(http.MethodPost, "/user_interaction", bytes.NewBuffer(jsonValue))
		context.Request.Header.Set("Content-Type", "application/json")

		handler.SaveUserInteraction(context)

		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Contains(t, writer.Body.String(), "User interactions saved successfully")
	})

	t.Run("Invalid Input", func(t *testing.T) {
		t.Parallel()

		writer := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(writer)
		handler := handlers.NewUserInteractionHandler(services_mock.UserInteractionServiceMock{IsOk: true})

		// Invalid JSON
		context.Request = httptest.NewRequest(http.MethodPost, "/user_interaction", bytes.NewBuffer([]byte(`invalid`)))
		context.Request.Header.Set("Content-Type", "application/json")

		handler.SaveUserInteraction(context)

		assert.Equal(t, http.StatusBadRequest, writer.Code)
		assert.Contains(t, writer.Body.String(), "Error parsing request")
	})
	t.Run("Fail on saving", func(t *testing.T) {
		t.Parallel()

		writer := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(writer)
		handler := handlers.NewUserInteractionHandler(services_mock.UserInteractionServiceMock{IsOk: false})
		input := []dtos.UserInteraction{
			{
				UserID:               "1",
				Action:               dtos.AddToCart,
				ProductSKU:           "123",
				InteractionTimestamp: time.Now(),
				InteractionDuration:  10,
			},
		}
		jsonValue, _ := json.Marshal(input)

		// Invalid JSON
		context.Request = httptest.NewRequest(http.MethodPost, "/user_interaction", bytes.NewBuffer((jsonValue)))
		context.Request.Header.Set("Content-Type", "application/json")

		handler.SaveUserInteraction(context)

		assert.Equal(t, http.StatusInternalServerError, writer.Code)
		assert.Contains(t, writer.Body.String(), "Error saving user interactions")
	})
}

func TestGetUserInteraction(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	t.Run("Get User Interaction", func(t *testing.T) {
		t.Parallel()

		writer := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(writer)
		handler := handlers.NewUserInteractionHandler(services_mock.UserInteractionServiceMock{IsOk: true})

		context.Params = gin.Params{gin.Param{Key: "user_id", Value: "1"}}
		context.Request = httptest.NewRequest(http.MethodGet, "/user_interaction/1", nil)
		context.Request.Header.Set("Content-Type", "application/json")

		handler.GetUserInteraction(context)

		assert.Equal(t, http.StatusOK, writer.Code)
	})
	t.Run("Should fail to get user Interaction", func(t *testing.T) {
		t.Parallel()
		writer := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(writer)
		handler := handlers.NewUserInteractionHandler(services_mock.UserInteractionServiceMock{IsOk: false})

		context.Params = gin.Params{gin.Param{Key: "user_id", Value: "1"}}
		context.Request = httptest.NewRequest(http.MethodGet, "/user_interaction/1", nil)
		context.Request.Header.Set("Content-Type", "application/json")

		handler.GetUserInteraction(context)

		assert.Equal(t, http.StatusNotFound, writer.Code)
	})
}
