package services_test

import (
	"testing"
	"time"

	resources_mock "yalo-api/mocks/resources"
	"yalo-api/pkg/handlers/dtos"
	"yalo-api/pkg/resources/interfaces"
	"yalo-api/pkg/services"
)

type UserInteractionTestSuit struct {
	name           string
	item           []dtos.UserInteraction
	dataRepository interfaces.UserInteractionDB
	want           func(t *testing.T, err error)
}

func TestUserInteractionService_Set(t *testing.T) {
	t.Parallel()

	tests := []UserInteractionTestSuit{
		{
			name: "Should fail on set User Interaction",
			item: []dtos.UserInteraction{
				{
					UserID:               "1",
					Action:               dtos.AddToCart,
					ProductSKU:           "123",
					InteractionTimestamp: time.Now(),
					InteractionDuration:  10,
				},
			},
			dataRepository: &resources_mock.UserInteractionDBMock{ErrorOnSet: true},
			want: func(t *testing.T, err error) {
				t.Helper()
				if err == nil {
					t.Errorf("want error, got nil")
				}
			},
		},
		{
			name: "Should works ok on set User Interaction",
			item: []dtos.UserInteraction{
				{
					UserID:               "1",
					Action:               dtos.AddToCart,
					ProductSKU:           "123",
					InteractionTimestamp: time.Now(),
					InteractionDuration:  10,
				},
			},
			dataRepository: &resources_mock.UserInteractionDBMock{ErrorOnSet: false},
			want: func(t *testing.T, err error) {
				t.Helper()
				if err != nil {
					t.Errorf("want nil, got error")
				}
			},
		},
	}
	for _, tt := range tests {
		testcase := tt

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			err := services.NewUserInteractionService(testcase.dataRepository).Save(testcase.item)

			testcase.want(t, err)
		})
	}
}

//nolint:funlen
func TestUserInteractionService_Get(t *testing.T) {
	t.Parallel()

	tests := []UserInteractionTestSuit{
		{
			name: "Should fail on Get User Interaction",
			item: []dtos.UserInteraction{
				{
					UserID:               "1",
					Action:               dtos.AddToCart,
					ProductSKU:           "123",
					InteractionTimestamp: time.Now(),
					InteractionDuration:  10,
				},
			},
			dataRepository: &resources_mock.UserInteractionDBMock{ErrorOnGetByUser: true},
			want: func(t *testing.T, err error) {
				t.Helper()
				if err == nil {
					t.Errorf("want error, got nil")
				}
			},
		},
		{
			name: "Should work with 2 User Interaction",
			item: []dtos.UserInteraction{
				{
					UserID:               "1",
					Action:               dtos.AddToCart,
					ProductSKU:           "123",
					InteractionTimestamp: time.Now(),
					InteractionDuration:  10,
				},
			},
			dataRepository: &resources_mock.UserInteractionDBMock{Return3: false},
			want: func(t *testing.T, err error) {
				t.Helper()
				if err != nil {
					t.Errorf("want nil, got error")
				}
			},
		},
		{
			name: "Should work with 3 User Interaction",
			item: []dtos.UserInteraction{
				{
					UserID:               "1",
					Action:               dtos.AddToCart,
					ProductSKU:           "123",
					InteractionTimestamp: time.Now(),
					InteractionDuration:  10,
				},
			},
			dataRepository: &resources_mock.UserInteractionDBMock{Return3: true},
			want: func(t *testing.T, err error) {
				t.Helper()
				if err != nil {
					t.Errorf("want nil, got error")
				}
			},
		},
	}
	for _, tt := range tests {
		testcase := tt

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			_, err := services.NewUserInteractionService(testcase.dataRepository).Get(testcase.item[0].UserID)

			testcase.want(t, err)
		})
	}
}
