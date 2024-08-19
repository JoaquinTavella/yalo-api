package repository_test

import (
	"testing"
	"time"
	"yalo-api/pkg/handlers/dtos"
	"yalo-api/pkg/resources/models"
	"yalo-api/pkg/resources/repository"
	"yalo-api/pkg/services"
)

type InMemoryDBTestSuite struct {
	name string
	item models.UserInteraction
	want func(t *testing.T, err bool, dataBase *repository.UserInteractionDB, item models.UserInteraction)
}

//nolint:funlen
func TestInMemoryDB(t *testing.T) {
	t.Parallel()

	tests := []InMemoryDBTestSuite{
		{
			name: "Should Set and Get User Interaction",
			item: models.UserInteraction{
				UserID:               "1",
				Action:               dtos.AddToCart,
				ProductSKU:           "123",
				InteractionTimestamp: time.Now(),
				InteractionDuration:  10,
			},
			want: func(t *testing.T, err bool, dataBase *repository.UserInteractionDB, item models.UserInteraction) {
				t.Helper()
				_, found := dataBase.Get(services.CreateKey(item.UserID, item.InteractionTimestamp, item.Action, item.ProductSKU))
				if !found {
					t.Errorf("want found, got not found")
				}
			},
		},
		{
			name: "Should Set and not found User Interaction",
			item: models.UserInteraction{
				UserID:               "1",
				Action:               dtos.AddToCart,
				ProductSKU:           "123",
				InteractionTimestamp: time.Now(),
				InteractionDuration:  10,
			},
			want: func(t *testing.T, err bool, dataBase *repository.UserInteractionDB, item models.UserInteraction) {
				t.Helper()
				_, found := dataBase.Get(services.CreateKey("123", item.InteractionTimestamp, item.Action, item.ProductSKU))
				if found {
					t.Errorf("want not found, got found")
				}
			},
		},
		{
			name: "Should Set delete and not found User Interaction",
			item: models.UserInteraction{
				UserID:               "1",
				Action:               dtos.AddToCart,
				ProductSKU:           "123",
				InteractionTimestamp: time.Now(),
				InteractionDuration:  10,
			},
			want: func(t *testing.T, err bool, dataBase *repository.UserInteractionDB, item models.UserInteraction) {
				t.Helper()
				key := services.CreateKey(item.UserID, item.InteractionTimestamp, item.Action, item.ProductSKU)
				dataBase.Delete(key)
				_, found := dataBase.Get(key)
				if found {
					t.Errorf("want not found, got found")
				}
			},
		},
		{
			name: "Should Set and get all User Interaction",
			item: models.UserInteraction{
				UserID:               "1",
				Action:               dtos.AddToCart,
				ProductSKU:           "123",
				InteractionTimestamp: time.Now(),
				InteractionDuration:  10,
			},
			want: func(t *testing.T, err bool, dataBase *repository.UserInteractionDB, item models.UserInteraction) {
				t.Helper()
				dataBase.Set(services.CreateKey("2", item.InteractionTimestamp, item.Action, item.ProductSKU), item)
				foundMap := dataBase.GetAll()
				foundMapLength := len(foundMap)
				if foundMapLength != 2 {
					t.Errorf("want 2, got %d", foundMapLength)
				}
			},
		},
		{
			name: "Should Set and get all User Interaction for 1 user",
			item: models.UserInteraction{
				UserID:               "1",
				Action:               dtos.AddToCart,
				ProductSKU:           "123",
				InteractionTimestamp: time.Now(),
				InteractionDuration:  10,
			},
			want: func(t *testing.T, err bool, dataBase *repository.UserInteractionDB, item models.UserInteraction) {
				t.Helper()
				secondItem := models.UserInteraction{
					UserID:               "2",
					Action:               dtos.AddToCart,
					ProductSKU:           "123",
					InteractionTimestamp: time.Now(),
					InteractionDuration:  10,
				}
				dataBase.Set(services.CreateKey("2", secondItem.InteractionTimestamp, secondItem.Action, secondItem.ProductSKU), secondItem)
				foundMap := dataBase.GetAllByUserID(item.UserID)
				foundMapLength := len(foundMap)
				if foundMapLength != 1 {
					t.Errorf("want 1, got %d", foundMapLength)
				}
			},
		},
	}

	for _, tt := range tests {
		testcase := tt

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			resource := repository.NewUserInteractionDB()

			res := resource.Set(services.CreateKey(testcase.item.UserID, testcase.item.InteractionTimestamp, testcase.item.Action, testcase.item.ProductSKU), testcase.item)

			testcase.want(t, res, resource, testcase.item)
		})
	}
}
