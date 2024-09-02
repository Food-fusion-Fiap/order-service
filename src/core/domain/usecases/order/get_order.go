package usecases

import (
	"github.com/Food-fusion-Fiap/order-service/src/adapters/gateways"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"log"
)

type GetOrderUseCase struct {
	OrderRepository gateways.OrderRepository
}

func (r *GetOrderUseCase) Execute(orderId string) (*entities.Order, error) {
	order, err := r.OrderRepository.FindById(orderId)
	if err != nil {
		log.Panic(err, "Error on GetOder")
		return nil, err
	}
	return order, err
}
