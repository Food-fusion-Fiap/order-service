package usecases

import (
	gateways2 "github.com/Food-fusion-Fiap/order-service/src/adapters/gateways"
)

type DeleteProductUsecase struct {
	repository gateways2.ProductRepository
}

func BuildDeleteProductUsecase(repository gateways2.ProductRepository) *DeleteProductUsecase {
	return &DeleteProductUsecase{repository: repository}
}

func (p *DeleteProductUsecase) Execute(id uint) error {
	return p.repository.DeleteById(id)
}
