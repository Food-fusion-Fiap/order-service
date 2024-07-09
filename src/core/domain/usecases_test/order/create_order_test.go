package order

import (
	usecases "github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases/order"
	"math/rand"
	"testing"

	"github.com/Food-fusion-Fiap/order-service/src/adapters/gateways"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/dtos"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"github.com/stretchr/testify/assert"
)

type mockCustomerRepository struct {
	gateways.CustomerRepository
	mockFindFirstById func(uint) (*entities.Customer, error)
}

func (m *mockCustomerRepository) FindFirstById(id uint) (*entities.Customer, error) {
	return m.mockFindFirstById(id)
}

type mockProductRepository struct {
	gateways.ProductRepository
	mockFindByIds func([]uint) ([]entities.Product, error)
}

func (m *mockProductRepository) FindByIds(ids []uint) ([]entities.Product, error) {
	return m.mockFindByIds(ids)
}

func TestCreateOrderUsecase_Execute(t *testing.T) {
	mockOrderRepo := &MockOrderRepository{}
	mockCustomerRepo := &mockCustomerRepository{}
	mockProductRepo := &mockProductRepository{}

	usecase := usecases.CreateOrderUsecase{
		OrderRepository:    mockOrderRepo,
		CustomerRepository: mockCustomerRepo,
		ProductRepository:  mockProductRepo,
	}

	t.Run("valid input and arbitrary customer", func(t *testing.T) {
		mockOrderRepo.mockCreate = func(order *entities.Order) (*entities.Order, error) {
			return &entities.Order{}, nil
		}

		mockCustomerRepo.mockFindFirstById = func(id uint) (*entities.Customer, error) {
			return &entities.Customer{ID: id}, nil
		}

		mockProductRepo.mockFindByIds = func(ids []uint) ([]entities.Product, error) {
			return []entities.Product{{ID: 1}, {ID: 2}}, nil
		}

		inputDto := dtos.CreateOrderDto{
			CustomerId: uint(rand.Intn(100)),
			Products: []dtos.ProductInsideOrder{
				{Id: 1, Quantity: 2, Observation: "test"},
				{Id: 2, Quantity: 1, Observation: "test"},
			},
		}

		_, err := usecase.Execute(inputDto)
		assert.NoError(t, err)
	})

	t.Run("invalid product", func(t *testing.T) {
		mockCustomerRepo.mockFindFirstById = func(id uint) (*entities.Customer, error) {
			return &entities.Customer{ID: id}, nil
		}

		mockProductRepo.mockFindByIds = func(ids []uint) ([]entities.Product, error) {
			return []entities.Product{}, nil
		}

		inputDto := dtos.CreateOrderDto{
			CustomerId: 1,
			Products: []dtos.ProductInsideOrder{
				{Id: 1, Quantity: 2, Observation: "test"},
				{Id: 2, Quantity: 1, Observation: "test"},
			},
		}

		_, err := usecase.Execute(inputDto)
		assert.EqualError(t, err, "algum dos produtos não foi encontrado")
	})
}
