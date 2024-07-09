package product

import (
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"github.com/Food-fusion-Fiap/order-service/src/infra/db/repositories"
	"github.com/stretchr/testify/mock"
)

// MockProductCategoryRepository mocks ProductCategoryRepository for testing
type MockProductCategoryRepository struct {
	repositories.ProductCategoryRepository
	mock.Mock
}

func (m *MockProductCategoryRepository) FindById(id uint) (*entities.ProductCategory, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.ProductCategory), args.Error(1)
}

func (m *MockProductCategoryRepository) IsExistingProductCategory() bool {
	return true // Simulating that the category exists
}
