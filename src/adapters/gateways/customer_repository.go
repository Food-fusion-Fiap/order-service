package gateways

import (
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
)

type CustomerRepository interface {
	Create(customer *entities.Customer) (*entities.Customer, error)
	List(customer *entities.Customer) ([]entities.Customer, error)
	FindFirstById(id uint) (*entities.Customer, error)
}
