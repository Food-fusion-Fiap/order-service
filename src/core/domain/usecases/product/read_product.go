package usecases

import (
	"github.com/Food-fusion-Fiap/order-service/src/adapters/gateways"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
)

type ReadProductUsecase struct {
	repository gateways.ProductRepository
}

func BuildReadProductUsecase(repository gateways.ProductRepository) *ReadProductUsecase {
	return &ReadProductUsecase{repository: repository}
}

func (p ReadProductUsecase) Execute(id uint) (*entities.Product, error) {
	return p.repository.FindById(id)
}
