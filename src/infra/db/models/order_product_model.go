package models

import "github.com/Food-fusion-Fiap/order-service/src/infra/db/mongo_driver/models"

type OrderProduct struct {
	Quantity    int            `bson:"quantity"`
	Observation string         `bson:"observation"`
	Product     models.Product `bson:"product"`
}
