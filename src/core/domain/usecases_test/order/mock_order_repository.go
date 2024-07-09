package order

import (
	"github.com/Food-fusion-Fiap/order-service/src/adapters/gateways"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
	gateways.OrderRepository
	mockCreate func(*entities.Order) (*entities.Order, error)
}

func (m *MockOrderRepository) Create(order *entities.Order) (*entities.Order, error) {
	return m.mockCreate(order)
}

func (m *MockOrderRepository) FindById(orderId string) (*entities.Order, error) {
	args := m.Called(orderId)
	return args.Get(0).(*entities.Order), args.Error(1)
}

func (m *MockOrderRepository) Update(order *entities.Order) (*entities.Order, error) {
	args := m.Called(order)
	return args.Get(0).(*entities.Order), args.Error(1)
}
