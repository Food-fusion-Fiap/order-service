package models

import (
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"github.com/Food-fusion-Fiap/order-service/src/utils"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID    uint
	Products      []OrderProduct
	Status        string
	PaymentStatus string
}

func (o Order) ToDomain(products []entities.ProductInsideOrder) entities.Order {
	return entities.Order{
		ID:            o.ID,
		CustomerID:    o.CustomerID,
		Products:      products,
		Status:        o.Status,
		PaymentStatus: o.PaymentStatus,
		CreatedAt:     o.CreatedAt.Format(utils.CompleteEnglishDateFormat),
		UpdatedAt:     o.UpdatedAt.Format(utils.CompleteEnglishDateFormat),
	}
}
