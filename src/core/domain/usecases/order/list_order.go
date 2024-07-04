package usecases

import (
	"github.com/Food-fusion-Fiap/order-service/src/adapters/gateways"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
)

type ListOrderUsecase struct {
	OrderRepository gateways.OrderRepository
}

func (r *ListOrderUsecase) Execute(sortBy string, orderBy string, status string) ([]entities.Order, error) {
	order, err := r.OrderRepository.List(sortBy, orderBy, status)

	return order, err
}
