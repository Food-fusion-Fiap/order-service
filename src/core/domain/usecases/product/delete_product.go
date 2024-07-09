package usecases

import (
	gateways2 "github.com/Food-fusion-Fiap/order-service/src/adapters/gateways"
)

type DeleteProductUsecase struct {
	repository      gateways2.ProductRepository
	orderRepository gateways2.OrderRepository
}

func BuildDeleteProductUsecase(repository gateways2.ProductRepository, orderRepository gateways2.OrderRepository) *DeleteProductUsecase {
	return &DeleteProductUsecase{repository: repository, orderRepository: orderRepository}
}

func (p *DeleteProductUsecase) Execute(id uint) error {
	return p.repository.DeleteById(id)
}
