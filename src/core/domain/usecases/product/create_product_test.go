package usecases

import (
	"errors"
	productDtos "github.com/Food-fusion-Fiap/order-service/src/core/domain/dtos/product"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	utils2 "github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateProductUsecase_Execute(t *testing.T) {

	mockProductRepo := new(utils2.MockProductRepository)
	mockProductCategoryRepo := new(utils2.MockProductCategoryRepository)

	emptyProductCategory := entities.ProductCategory{}
	productCategory := entities.ProductCategory{ID: 1, Description: "desc"}

	usecase := BuildCreateProductUsecase(mockProductRepo, mockProductCategoryRepo)

	tests := []struct {
		name                    string
		inputDto                productDtos.PersistProductDto
		expectedProduct         *entities.Product
		expectedProductCategory *entities.ProductCategory
		productRepoError        error
		categoryNotFound        bool
		categoryRepoError       error
		expectedError           error
	}{
		{
			name: "Successful Product Creation",
			inputDto: productDtos.PersistProductDto{
				Name:        "Test Product",
				Price:       10.99,
				Description: "Test Description",
				CategoryID:  1,
			},
			expectedProduct: &entities.Product{
				Name:        "Test Product",
				Price:       10.99,
				Description: "Test Description",
				CategoryId:  1,
			},
			expectedProductCategory: &productCategory,
			productRepoError:        nil,
			categoryNotFound:        false,
			categoryRepoError:       nil,
			expectedError:           nil,
		},
		{
			name: "Product Category Not Found",
			inputDto: productDtos.PersistProductDto{
				Name:        "Test Product",
				Price:       10.99,
				Description: "Test Description",
				CategoryID:  2,
			},
			expectedProduct: &entities.Product{
				Name:        "Test Product",
				Price:       10.99,
				Description: "Test Description",
				CategoryId:  2,
			},
			expectedProductCategory: &emptyProductCategory,
			productRepoError:        nil,
			categoryNotFound:        true,
			categoryRepoError:       nil,
			expectedError:           errors.New("product category doesn't exist"),
		},
		{
			name: "Error Finding Product Category",
			inputDto: productDtos.PersistProductDto{
				Name:        "Test Product",
				Price:       10.99,
				Description: "Test Description",
				CategoryID:  3,
			},
			expectedProduct:   nil,
			productRepoError:  nil,
			categoryNotFound:  false,
			categoryRepoError: errors.New("ProductCategory repository error"),
			expectedError:     errors.New("ProductCategory repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockProductCategoryRepo.On("FindById", tt.inputDto.CategoryID).
				Return(tt.expectedProductCategory, tt.categoryRepoError).
				Once()

			if !tt.categoryNotFound && tt.categoryRepoError == nil {
				mockProductRepo.On("Create", mock.AnythingOfType("*entities.Product")).
					Return(tt.expectedProduct, nil).
					Once()
			}

			product, err := usecase.Execute(tt.inputDto)

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedProduct, product)
			}

			mockProductRepo.AssertExpectations(t)
			mockProductCategoryRepo.AssertExpectations(t)
		})
	}
}
