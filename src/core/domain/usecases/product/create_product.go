package usecases

import (
	"errors"
	"github.com/Food-fusion-Fiap/order-service/src/adapters/gateways"
	productDtos "github.com/Food-fusion-Fiap/order-service/src/core/domain/dtos/product"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"log"
)

type CreateProductUsecase struct {
	repository                gateways.ProductRepository
	productCategoryRepository gateways.ProductCategoryRepository
}

func BuildCreateProductUsecase(repository gateways.ProductRepository,
	productCategoryRepository gateways.ProductCategoryRepository) *CreateProductUsecase {
	return &CreateProductUsecase{repository: repository, productCategoryRepository: productCategoryRepository}
}

func (p *CreateProductUsecase) Execute(inputDto productDtos.PersistProductDto) (*entities.Product, error) {
	product := &entities.Product{
		Name:        inputDto.Name,
		Price:       inputDto.Price,
		Description: inputDto.Description,
		CategoryId:  inputDto.CategoryID,
	}

	category, err := p.productCategoryRepository.FindById(product.CategoryId)

	if err != nil {
		return nil, errors.New("ProductCategory repository error")
	}

	if !category.IsExistingProductCategory() {
		log.Println("Product category doesn't exist! - productCategoryId=", product.CategoryId)
		return product, errors.New("product category doesn't exist")
	}

	return p.repository.Create(product)
}
