package main

import (
	"github.com/Food-fusion-Fiap/order-service/src/infra/db/gorm"
	"github.com/Food-fusion-Fiap/order-service/src/infra/web/routes"
)

func main() {
	gorm.ConnectDB()
	routes.HandleRequests()
}
