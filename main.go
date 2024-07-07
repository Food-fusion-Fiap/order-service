package main

import (
	"github.com/Food-fusion-Fiap/order-service/src/infra/db/gorm"
	"github.com/Food-fusion-Fiap/order-service/src/infra/db/mongo_driver"
	"github.com/Food-fusion-Fiap/order-service/src/infra/web/routes"
)

func main() {
	gorm.ConnectDB()
	mongo_driver.ConnectDB()
	routes.HandleRequests()
}
