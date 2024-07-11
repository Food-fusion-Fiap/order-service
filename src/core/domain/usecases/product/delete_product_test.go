package usecases

import (
	"errors"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteProductUsecase_Execute_Success(t *testing.T) {
	// Arrange
	repo := new(utils.MockProductRepository)
	usecase := BuildDeleteProductUsecase(repo)

	repo.On("DeleteById", uint(1)).Return(nil)

	// Act
	err := usecase.Execute(1)

	// Assert
	assert.Nil(t, err)

	repo.AssertExpectations(t)
}

func TestDeleteProductUsecase_Execute_Error(t *testing.T) {
	// Arrange
	repo := new(utils.MockProductRepository)
	usecase := BuildDeleteProductUsecase(repo)

	repo.On("DeleteById", uint(1)).Return(errors.New("error deleting product"))

	// Act
	err := usecase.Execute(1)

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, "error deleting product", err.Error())

	repo.AssertExpectations(t)
}
