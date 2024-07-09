package models

import (
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"github.com/Food-fusion-Fiap/order-service/src/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID    int64              `bson:"customer_id"`
	Products      []OrderProduct     `bson:"products"`
	Status        string             `bson:"status"`
	PaymentStatus string             `bson:"payment_status"`
	CreatedAt     primitive.DateTime `bson:"created_at,omitempty"`
	UpdatedAt     primitive.DateTime `bson:"updated_at,omitempty"`
}

func (o Order) ToDomain(products []entities.ProductInsideOrder) *entities.Order {
	return &entities.Order{
		ID:            o.ID.Hex(),
		CustomerID:    o.CustomerID,
		Products:      products,
		Status:        o.Status,
		PaymentStatus: o.PaymentStatus,
		CreatedAt:     o.CreatedAt.Time().Format(utils.CompleteEnglishDateFormat),
		UpdatedAt:     o.UpdatedAt.Time().Format(utils.CompleteEnglishDateFormat),
	}
}
