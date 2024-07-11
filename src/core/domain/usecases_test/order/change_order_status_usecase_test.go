package order

import (
	"errors"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/enums"
	usecases "github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases/order"
	utils2 "github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases_test/utils"
	"github.com/Food-fusion-Fiap/order-service/src/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestChangeOrderStatusUsecase_Execute(t *testing.T) {
	// Define common test variables
	orderId := primitive.NewObjectID().Hex()

	tests := []struct {
		name           string
		initialOrder   *entities.Order
		changeToStatus string
		expectedOrder  *entities.Order
		expectedError  error
	}{
		{
			name: "Change to Received - Success",
			initialOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Created,
				PaymentStatus: enums.AwaitingPayment,
			},
			changeToStatus: enums.Received,
			expectedOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Received,
				PaymentStatus: enums.Paid,
				UpdatedAt:     time.Now().Format(utils.CompleteEnglishDateFormat),
			},
			expectedError: nil,
		},
		{
			name: "Change to Cancelled - Success",
			initialOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Created,
				PaymentStatus: enums.AwaitingPayment,
			},
			changeToStatus: enums.Cancelled,
			expectedOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Cancelled,
				PaymentStatus: enums.Cancelled,
				UpdatedAt:     time.Now().Format(utils.CompleteEnglishDateFormat),
			},
			expectedError: nil,
		},
		{
			name: "Change to Cancelled - Failure",
			initialOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Received,
				PaymentStatus: enums.Paid,
			},
			changeToStatus: enums.Cancelled,
			expectedOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Received,
				PaymentStatus: enums.Paid,
			},
			expectedError: errors.New("não é possível cancelar o pedido"),
		},
		{
			name: "Change to Preparation - Success",
			initialOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Received,
				PaymentStatus: enums.Paid,
			},
			changeToStatus: enums.Preparation,
			expectedOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Preparation,
				PaymentStatus: enums.Paid,
				UpdatedAt:     time.Now().Format(utils.CompleteEnglishDateFormat),
			},
			expectedError: nil,
		},
		{
			name: "Change to Preparation - Failure",
			initialOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Created,
				PaymentStatus: enums.AwaitingPayment,
			},
			changeToStatus: enums.Preparation,
			expectedOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Created,
				PaymentStatus: enums.AwaitingPayment,
			},
			expectedError: errors.New("não é possível mudar o pedido para Em preparação"),
		},
		{
			name: "Change to Ready - Success",
			initialOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Preparation,
				PaymentStatus: enums.Paid,
			},
			changeToStatus: enums.Ready,
			expectedOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Ready,
				PaymentStatus: enums.Paid,
				UpdatedAt:     time.Now().Format(utils.CompleteEnglishDateFormat),
			},
			expectedError: nil,
		},
		{
			name: "Change to Ready - Failure",
			initialOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Received,
				PaymentStatus: enums.Paid,
			},
			changeToStatus: enums.Ready,
			expectedOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Received,
				PaymentStatus: enums.Paid,
			},
			expectedError: errors.New("não é possível mudar o pedido para Pronto"),
		},
		{
			name: "Change to Finished - Success",
			initialOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Ready,
				PaymentStatus: enums.Paid,
			},
			changeToStatus: enums.Finished,
			expectedOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Finished,
				PaymentStatus: enums.Paid,
				UpdatedAt:     time.Now().Format(utils.CompleteEnglishDateFormat),
			},
			expectedError: nil,
		},
		{
			name: "Change to Finished - Failure",
			initialOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Preparation,
				PaymentStatus: enums.Paid,
			},
			changeToStatus: enums.Finished,
			expectedOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Preparation,
				PaymentStatus: enums.Paid,
			},
			expectedError: errors.New("não é possível mudar o pedido para Finalizado"),
		},
		{
			name: "Unknown Status",
			initialOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Created,
				PaymentStatus: enums.AwaitingPayment,
			},
			changeToStatus: "unknown",
			expectedOrder: &entities.Order{
				ID:            orderId,
				Status:        enums.Created,
				PaymentStatus: enums.AwaitingPayment,
			},
			expectedError: errors.New("status desconhecido"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(utils2.MockOrderRepository)
			usecase := usecases.ChangeOrderStatusUsecase{OrderRepository: mockRepo}

			mockRepo.On("FindById", orderId).Return(tt.initialOrder, nil)
			if tt.expectedError == nil {
				mockRepo.On("Update", mock.AnythingOfType("*entities.Order")).Return(tt.expectedOrder, nil)
			}

			order, err := usecase.Execute(orderId, tt.changeToStatus)

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOrder, order)
				assert.Equal(t, tt.expectedOrder.UpdatedAt, order.UpdatedAt)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
