package order

import (
	"errors"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/enums"
	usecases "github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases/order"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases_test/utils"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/Food-fusion-Fiap/order-service/src/core/domain/dtos"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrderUsecase_Execute_Success(t *testing.T) {
	// Arrange
	orderRepo := new(utils.MockOrderRepository)
	productRepo := new(utils.MockProductRepository)
	usecase := usecases.CreateOrderUsecase{
		OrderRepository:    orderRepo,
		CustomerRepository: nil,
		ProductRepository:  productRepo,
	}

	inputDto := dtos.CreateOrderDto{
		CustomerId: 1,
		Products:   []dtos.ProductInsideOrder{{Id: 1, Quantity: 2, Observation: "Test observation"}},
	}
	expectedOrder := &entities.Order{
		Status:        enums.Created,
		PaymentStatus: enums.AwaitingPayment,
		CustomerID:    1,
		Products: []entities.ProductInsideOrder{
			{Quantity: 2, Observation: "Test observation", Product: entities.Product{ID: 1}},
		},
	}

	product := &entities.Product{ID: 1}

	productRepo.On("FindByIds", []uint{1}).Return([]entities.Product{*product}, nil).Twice()
	orderRepo.On("Create", mock.AnythingOfType("*entities.Order")).Return(expectedOrder, nil)

	// Act
	result, err := usecase.Execute(inputDto)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedOrder, result)

	productRepo.AssertExpectations(t)
	orderRepo.AssertExpectations(t)
}

func TestCreateOrderUsecase_Execute_ProductNotFound(t *testing.T) {
	// Arrange
	orderRepo := new(utils.MockOrderRepository)
	productRepo := new(utils.MockProductRepository)
	usecase := usecases.CreateOrderUsecase{
		OrderRepository:    orderRepo,
		CustomerRepository: nil,
		ProductRepository:  productRepo,
	}

	inputDto := dtos.CreateOrderDto{
		CustomerId: 1,
		Products:   []dtos.ProductInsideOrder{{Id: 1, Quantity: 2, Observation: "Test observation"}},
	}

	productRepo.On("FindByIds", []uint{1}).Return([]entities.Product{}, errors.New("algum dos produtos não foi encontrado"))

	// Act
	result, err := usecase.Execute(inputDto)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "algum dos produtos não foi encontrado", err.Error())

	productRepo.AssertExpectations(t)
}

func TestCreateOrderUsecase_Execute_OrderCreationError(t *testing.T) {
	//Arrange
	orderRepo := new(utils.MockOrderRepository)
	productRepo := new(utils.MockProductRepository)
	usecase := usecases.CreateOrderUsecase{
		OrderRepository:    orderRepo,
		CustomerRepository: nil,
		ProductRepository:  productRepo,
	}

	inputDto := dtos.CreateOrderDto{
		CustomerId: 1,
		Products:   []dtos.ProductInsideOrder{{Id: 1, Quantity: 2, Observation: "Test observation"}},
	}

	product := &entities.Product{ID: 1}

	productRepo.On("FindByIds", []uint{1}).Return([]entities.Product{*product}, nil)
	orderRepo.On("Create", mock.AnythingOfType("*entities.Order")).Return(&entities.Order{}, errors.New("error creating order"))

	//Act
	result, err := usecase.Execute(inputDto)

	//Assert
	assert.NotNil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "error creating order", err.Error())

	productRepo.AssertExpectations(t)
	orderRepo.AssertExpectations(t)
}
