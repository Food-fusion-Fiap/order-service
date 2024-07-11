package gateways

import (
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
)

type OrderRepository interface {
	List(sortBy string, orderBy string, status string) ([]entities.Order, error)
	FindById(orderId string) (*entities.Order, error)
	Update(*entities.Order) (*entities.Order, error)
	Create(order *entities.Order) (*entities.Order, error)
	GetDescOrder() string
	GetAscOrder() string
	GetCreatedAtFieldName() string
}
