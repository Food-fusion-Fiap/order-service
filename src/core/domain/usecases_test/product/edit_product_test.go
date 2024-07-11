package product

import (
	"errors"
	dtosProd "github.com/Food-fusion-Fiap/order-service/src/core/domain/dtos/product"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	usecases "github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases/product"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases_test/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestExecute_ExistingProduct(t *testing.T) {

	repo := new(utils.MockProductRepository)
	usecase := usecases.BuildEditProductUsecase(repo)

	inputDto := dtosProd.PersistProductDto{
		ID:          123,
		Name:        "New Product Name",
		Price:       100.0,
		Description: "New Description",
		CategoryID:  321,
	}

	existingProduct := &entities.Product{ID: inputDto.ID,
		Name:  "Old Product Name",
		Price: 50.0, Description: "Old Description",
		CategoryId: 432}
	repo.On("FindById", inputDto.ID).Return(existingProduct, nil)
	repo.On("Edit", mock.AnythingOfType("*entities.Product")).Return(existingProduct, nil)

	result, err := usecase.Execute(inputDto)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, inputDto.Name, result.Name)
	assert.Equal(t, inputDto.Price, result.Price)
	assert.Equal(t, inputDto.Description, result.Description)
	assert.Equal(t, inputDto.CategoryID, result.CategoryId)

	repo.AssertExpectations(t)
}

func TestExecute_NonExistingProduct(t *testing.T) {
	repo := new(utils.MockProductRepository)
	usecase := usecases.BuildEditProductUsecase(repo)

	inputDto := dtosProd.PersistProductDto{
		ID:          321,
		Name:        "New Product Name",
		Price:       100.0,
		Description: "New Description",
		CategoryID:  123,
	}

	repo.On("FindById", inputDto.ID).Return(&entities.Product{}, errors.New("product not found"))

	result, err := usecase.Execute(inputDto)

	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error())
	assert.Equal(t, &entities.Product{}, result)

	repo.AssertExpectations(t)
}
