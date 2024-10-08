package usecases

import (
	"github.com/Food-fusion-Fiap/order-service/src/adapters/gateways"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
)

type ListProductUsecase struct {
	repository gateways.ProductRepository
}

func BuildListProductUsecase(repository gateways.ProductRepository) *ListProductUsecase {
	return &ListProductUsecase{repository: repository}
}

func (p ListProductUsecase) Execute(categoryId uint) ([]entities.Product, error) {

	if categoryId == 0 {
		return p.repository.FindAll()
	}

	return p.repository.FindByCategoryId(categoryId)
}
