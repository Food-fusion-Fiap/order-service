package usecases

import (
	"github.com/Food-fusion-Fiap/order-service/src/adapters/gateways"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/enums"
	"slices"
)

type ListOngoingOrdersUsecase struct {
	OrderRepository gateways.OrderRepository
}

func (r *ListOngoingOrdersUsecase) Execute() ([]entities.Order, error) {
	var createdAt = r.OrderRepository.GetCreatedAtFieldName()
	var ascOrder = r.OrderRepository.GetAscOrder()

	ready, err := r.OrderRepository.List(createdAt, ascOrder, enums.Ready)
	preparation, err := r.OrderRepository.List(createdAt, ascOrder, enums.Preparation)
	received, err := r.OrderRepository.List(createdAt, ascOrder, enums.Received)

	result := slices.Concat(ready, preparation, received)

	return result, err
}
