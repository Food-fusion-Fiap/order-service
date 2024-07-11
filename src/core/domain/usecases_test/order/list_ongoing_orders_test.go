package order

import (
	"errors"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/enums"
	usecases "github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases/order"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListOngoingOrdersUsecase_Execute(t *testing.T) {
	// Define mock data
	mockOrders := []entities.Order{
		{ID: "1", Status: enums.Ready},
		{ID: "2", Status: enums.Preparation},
		{ID: "3", Status: enums.Received},
	}

	// Define test cases
	tests := []struct {
		name             string
		mockListResponse map[string][]entities.Order
		expectedResult   []entities.Order
		expectedError    error
	}{
		{
			name: "Successful Fetch",
			mockListResponse: map[string][]entities.Order{
				enums.Ready:       {mockOrders[0]},
				enums.Preparation: {mockOrders[1]},
				enums.Received:    {mockOrders[2]},
			},
			expectedResult: mockOrders,
			expectedError:  nil,
		},
		{
			name: "Repository Error",
			mockListResponse: map[string][]entities.Order{
				enums.Ready:       nil,
				enums.Preparation: nil,
				enums.Received:    nil,
			},
			expectedResult: nil,
			expectedError:  errors.New("error fetching orders"),
		},
	}

	for _, tt := range tests {
		mockOrderRepo := new(MockOrderRepository)
		usecase := new(usecases.ListOngoingOrdersUsecase)
		usecase.OrderRepository = mockOrderRepo

		t.Run(tt.name, func(t *testing.T) {
			// Set up mocks
			for status, orders := range tt.mockListResponse {
				mockOrderRepo.On("List", "created_at", "ASC", status).Return(orders, tt.expectedError).Once()
			}

			// Execute the use case
			result, err := usecase.Execute()

			// Verify results
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}

			// Assert that all expected repository calls were made
			mockOrderRepo.AssertExpectations(t)
		})
	}
}
