package product

import (
	"errors"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	usecasesProduct "github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases/product"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadProductUsecase_Execute_ProductFound(t *testing.T) {
	repo := new(MockProductRepository)
	usecase := usecasesProduct.BuildReadProductUsecase(repo)

	expectedProduct := &entities.Product{
		ID:          1,
		Name:        "Test Product",
		Price:       100.0,
		Description: "Test Description",
		CategoryId:  123,
	}

	repo.On("FindById", expectedProduct.ID).Return(expectedProduct, nil)

	result, err := usecase.Execute(expectedProduct.ID)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedProduct, result)

	repo.AssertExpectations(t)
}

func TestReadProductUsecase_Execute_ProductNotFound(t *testing.T) {
	repo := new(MockProductRepository)
	usecase := usecasesProduct.BuildReadProductUsecase(repo)

	var id uint = 1

	repo.On("FindById", id).Return(&entities.Product{}, errors.New("product not found"))

	result, err := usecase.Execute(id)

	assert.NotNil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, err.Error(), err.Error())

	repo.AssertExpectations(t)
}
