package usecases

import (
	"errors"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListProductUsecase_Execute_NoCategoryId(t *testing.T) {
	// Arrange
	repo := new(utils.MockProductRepository)
	usecase := BuildListProductUsecase(repo)

	expectedProducts := []entities.Product{
		{ID: 1, Name: "Product 1", Price: 100.0, Description: "Description 1", CategoryId: 456},
		{ID: 2, Name: "Product 2", Price: 200.0, Description: "Description 2", CategoryId: 123},
	}
	repo.On("FindAll").Return(expectedProducts, nil)

	// Act
	result, err := usecase.Execute(0)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedProducts, result)

	repo.AssertExpectations(t)
}

func TestListProductUsecase_Execute_WithCategoryId(t *testing.T) {
	// Arrange
	repo := new(utils.MockProductRepository)
	usecase := BuildListProductUsecase(repo)

	expectedProducts := []entities.Product{
		{ID: 1, Name: "Product 1", Price: 100.0, Description: "Description 1", CategoryId: 1},
	}
	repo.On("FindByCategoryId", uint(1)).Return(expectedProducts, nil)

	// Act
	result, err := usecase.Execute(1)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedProducts, result)

	repo.AssertExpectations(t)
}

func TestListProductUsecase_Execute_FindAllError(t *testing.T) {
	// Arrange
	repo := new(utils.MockProductRepository)
	usecase := BuildListProductUsecase(repo)

	repo.On("FindAll").Return([]entities.Product{}, errors.New("error finding all products"))

	// Act
	result, err := usecase.Execute(0)

	// Assert
	assert.NotNil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "error finding all products", err.Error())

	repo.AssertExpectations(t)
}

func TestListProductUsecase_Execute_FindByCategoryIdError(t *testing.T) {
	// Arrange
	repo := new(utils.MockProductRepository)
	usecase := BuildListProductUsecase(repo)

	repo.On("FindByCategoryId", uint(1)).Return([]entities.Product{}, errors.New("error finding products by category ID"))

	// Act
	result, err := usecase.Execute(1)

	// Assert
	assert.NotNil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "error finding products by category ID", err.Error())

	repo.AssertExpectations(t)
}
