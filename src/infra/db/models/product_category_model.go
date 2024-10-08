package models

import (
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"gorm.io/gorm"
)

type ProductCategory struct {
	gorm.Model
	Description string
}

func (pc ProductCategory) ToDomain() entities.ProductCategory {
	return entities.ProductCategory{
		ID:          pc.ID,
		Description: pc.Description,
	}
}
