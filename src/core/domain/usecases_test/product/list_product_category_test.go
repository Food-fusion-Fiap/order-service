package product

import (
	"errors"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	usecases "github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases/product"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListProductCategoryUsecase_Execute_Success(t *testing.T) {
	// Arrange
	repo := new(MockProductCategoryRepository)
	usecase := usecases.BuildListProductCategoryUsecase(repo)

	expectedCategories := []entities.ProductCategory{
		{ID: 1, Description: "Category 1"},
		{ID: 2, Description: "Category 2"},
	}
	repo.On("FindAll").Return(expectedCategories, nil)

	// Act
	result, err := usecase.Execute()

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedCategories, result)

	repo.AssertExpectations(t)
}

func TestListProductCategoryUsecase_Execute_Error(t *testing.T) {
	// Arrange
	repo := new(MockProductCategoryRepository)
	usecase := usecases.BuildListProductCategoryUsecase(repo)

	repo.On("FindAll").Return([]entities.ProductCategory{}, errors.New("error finding product categories"))

	// Act
	result, err := usecase.Execute()

	// Assert
	assert.NotNil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "error finding product categories", err.Error())

	repo.AssertExpectations(t)
}
