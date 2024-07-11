package gateways

import "github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"

type ProductCategoryRepository interface {
	FindAll() ([]entities.ProductCategory, error)
	FindById(id uint) (*entities.ProductCategory, error)
}
