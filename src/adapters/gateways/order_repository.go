package gateways

import (
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
)

type OrderRepository interface {
	List(sortBy string, orderBy string, status string) ([]entities.Order, error)
	FindById(orderId uint) *entities.Order
	Update(*entities.Order)
	Create(order *entities.Order) (*entities.Order, error)
	ExistsOrderProduct(productId uint) bool
	GetDescOrder() string
	GetAscOrder() string
	GetCreatedAtFieldName() string
}
