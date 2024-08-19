package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
	"yalo-api/internal/handlers/dtos"
	handlers "yalo-api/internal/handlers/users_interactions"
	services_mock "yalo-api/mocks/services"
)

//nolint:funlen
func TestSaveUserInteraction(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Parallel()

	t.Run("Should Fail because userID is missing", func(t *testing.T) {
		t.Parallel()

		writer := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(writer)

		context.Set("logger", zerolog.New(os.Stdout))

		handler := handlers.NewUserInteractionHandler(services_mock.UserInteractionServiceMock{IsOk: true})

		input := dtos.UserInteractionRequest{UserInteraction: []dtos.UserInteraction{
			{
				Action:               dtos.AddToCart,
				ProductSKU:           "123",
				InteractionTimestamp: time.Now(),
				InteractionDuration:  10,
			},
		},
		}
		jsonValue, _ := json.Marshal(input)

		context.Request = httptest.NewRequest(http.MethodPost, "/user_interaction", bytes.NewBuffer(jsonValue))
		context.Request.Header.Set("Content-Type", "application/json")

		handler.SaveUserInteraction(context)

		assert.Equal(t, http.StatusBadRequest, writer.Code)
	})

	t.Run("Should Fail because ProductSKU is missing", func(t *testing.T) {
		t.Parallel()

		writer := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(writer)

		context.Set("logger", zerolog.New(os.Stdout))

		handler := handlers.NewUserInteractionHandler(services_mock.UserInteractionServiceMock{IsOk: true})

		// Sample input
		input := dtos.UserInteractionRequest{UserInteraction: []dtos.UserInteraction{
			{
				UserID:               "1",
				Action:               dtos.AddToCart,
				ProductSKU:           "",
				InteractionTimestamp: time.Now(),
				InteractionDuration:  10,
			},
		},
		}
		jsonValue, _ := json.Marshal(input)

		context.Request = httptest.NewRequest(http.MethodPost, "/user_interaction", bytes.NewBuffer(jsonValue))
		context.Request.Header.Set("Content-Type", "application/json")

		handler.SaveUserInteraction(context)

		assert.Equal(t, http.StatusBadRequest, writer.Code)
	})

	t.Run("Should Fail because TIme is missing", func(t *testing.T) {
		t.Parallel()

		writer := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(writer)

		context.Set("logger", zerolog.New(os.Stdout))

		handler := handlers.NewUserInteractionHandler(services_mock.UserInteractionServiceMock{IsOk: true})

		// Sample input
		input := dtos.UserInteractionRequest{UserInteraction: []dtos.UserInteraction{
			{
				UserID:              "1",
				Action:              dtos.AddToCart,
				ProductSKU:          "234",
				InteractionDuration: 10,
			},
		},
		}
		jsonValue, _ := json.Marshal(input)

		context.Request = httptest.NewRequest(http.MethodPost, "/user_interaction", bytes.NewBuffer(jsonValue))
		context.Request.Header.Set("Content-Type", "application/json")

		handler.SaveUserInteraction(context)

		assert.Equal(t, http.StatusBadRequest, writer.Code)
	})

	t.Run("Should Fail because Action is wrong", func(t *testing.T) {
		t.Parallel()

		writer := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(writer)

		context.Set("logger", zerolog.New(os.Stdout))

		handler := handlers.NewUserInteractionHandler(services_mock.UserInteractionServiceMock{IsOk: true})

		// Sample input
		input := dtos.UserInteractionRequest{UserInteraction: []dtos.UserInteraction{
			{
				UserID:               "1",
				Action:               "231231",
				ProductSKU:           "234",
				InteractionTimestamp: time.Now(),
				InteractionDuration:  10,
			},
		},
		}
		jsonValue, _ := json.Marshal(input)

		context.Request = httptest.NewRequest(http.MethodPost, "/user_interaction", bytes.NewBuffer(jsonValue))
		context.Request.Header.Set("Content-Type", "application/json")

		handler.SaveUserInteraction(context)

		assert.Equal(t, http.StatusBadRequest, writer.Code)
	})

	t.Run("Should Fail because its an empty arr", func(t *testing.T) {
		t.Parallel()

		writer := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(writer)

		context.Set("logger", zerolog.New(os.Stdout))

		handler := handlers.NewUserInteractionHandler(services_mock.UserInteractionServiceMock{IsOk: true})

		// Sample input
		input := dtos.UserInteractionRequest{UserInteraction: []dtos.UserInteraction{}}
		jsonValue, _ := json.Marshal(input)

		context.Request = httptest.NewRequest(http.MethodPost, "/user_interaction", bytes.NewBuffer(jsonValue))
		context.Request.Header.Set("Content-Type", "application/json")

		handler.SaveUserInteraction(context)

		assert.Equal(t, http.StatusBadRequest, writer.Code)
	})

	t.Run("Successful Save", func(t *testing.T) {
		t.Parallel()

		writer := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(writer)

		context.Set("logger", zerolog.New(os.Stdout))

		handler := handlers.NewUserInteractionHandler(services_mock.UserInteractionServiceMock{IsOk: true})

		// Sample input
		input := dtos.UserInteractionRequest{UserInteraction: []dtos.UserInteraction{
			{
				UserID:               "1",
				Action:               dtos.AddToCart,
				ProductSKU:           "123",
				InteractionTimestamp: time.Now(),
				InteractionDuration:  10,
			},
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

		context.Set("logger", zerolog.New(os.Stdout))

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

		context.Set("logger", zerolog.New(os.Stdout))

		handler := handlers.NewUserInteractionHandler(services_mock.UserInteractionServiceMock{IsOk: false})
		input := dtos.UserInteractionRequest{UserInteraction: []dtos.UserInteraction{
			{
				UserID:               "1",
				Action:               dtos.AddToCart,
				ProductSKU:           "123",
				InteractionTimestamp: time.Now(),
				InteractionDuration:  10,
			},
		}}
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

		context.Set("logger", zerolog.New(os.Stdout))

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

		context.Set("logger", zerolog.New(os.Stdout))

		handler := handlers.NewUserInteractionHandler(services_mock.UserInteractionServiceMock{IsOk: false})

		context.Params = gin.Params{gin.Param{Key: "user_id", Value: "1"}}
		context.Request = httptest.NewRequest(http.MethodGet, "/user_interaction/1", nil)
		context.Request.Header.Set("Content-Type", "application/json")

		handler.GetUserInteraction(context)

		assert.Equal(t, http.StatusNotFound, writer.Code)
	})
}
