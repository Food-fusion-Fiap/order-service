package models

import (
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
)

type Product struct {
	ID          int64   `bson:"_id,omitempty"`
	Name        string  `bson:"name"`
	Price       float64 `bson:"price"`
	Description string  `bson:"description"`
	CategoryID  int64   `bson:"category_id"`
}

func (c Product) ToDomain() entities.Product {
	return entities.Product{
		ID:          uint(c.ID),
		Name:        c.Name,
		Price:       c.Price,
		Description: c.Description,
		CategoryId:  uint(c.CategoryID),
	}
}

func (p *Product) PatchFields(name string, price float64, description string, categoryId uint) {
	p.Name = name
	p.Price = price
	p.Description = description
	p.CategoryID = int64(categoryId)
}
