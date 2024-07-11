package product

import (
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"github.com/Food-fusion-Fiap/order-service/src/infra/db/repositories"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	repositories.ProductRepository
	mock.Mock
}

func (m *MockProductRepository) Create(product *entities.Product) (*entities.Product, error) {
	args := m.Called(product)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (m *MockProductRepository) FindById(id uint) (*entities.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (m *MockProductRepository) Edit(product *entities.Product) (*entities.Product, error) {
	args := m.Called(product)
	return args.Get(0).(*entities.Product), args.Error(1)
}
