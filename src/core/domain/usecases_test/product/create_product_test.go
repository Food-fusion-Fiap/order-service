package product

import (
	"errors"
	productDtos "github.com/Food-fusion-Fiap/order-service/src/core/domain/dtos/product"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	usecases "github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases/product"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateProductUsecase_Execute(t *testing.T) {
	// Initialize mocks
	mockProductRepo := new(MockProductRepository)
	mockProductCategoryRepo := new(MockProductCategoryRepository)

	emptyProductCategory := entities.ProductCategory{}
	productCategory := entities.ProductCategory{ID: 1, Description: "desc"}

	// Create instance of CreateProductUsecase
	usecase := usecases.BuildCreateProductUsecase(mockProductRepo, mockProductCategoryRepo)

	// Define test cases
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
				CategoryID:  2, // Assuming category ID 2 doesn't exist
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

	// Iterate over test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockProductCategoryRepo.On("FindById", tt.inputDto.CategoryID).
				Return(tt.expectedProductCategory, tt.categoryRepoError).
				Once() // Ensure this method is called once

			// Mock behavior of ProductRepository based on the scenario
			if !tt.categoryNotFound && tt.categoryRepoError == nil {
				mockProductRepo.On("Create", mock.AnythingOfType("*entities.Product")).
					Return(tt.expectedProduct, nil).
					Once() // Ensure this method is called once
			}

			// Execute the use case method
			product, err := usecase.Execute(tt.inputDto)

			// Assertions
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedProduct, product)
			}

			// Verify expectations
			mockProductRepo.AssertExpectations(t)
			mockProductCategoryRepo.AssertExpectations(t)
		})
	}
}
