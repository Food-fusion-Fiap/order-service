package routes

import (
	orderController "github.com/Food-fusion-Fiap/order-service/src/adapters/controllers/order"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine) {
	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("", orderController.CreateOrder)
		orderRoutes.GET("", orderController.ListOngoingOrders)
		orderRoutes.PATCH("", orderController.ChangeOrderStatus)
		orderRoutes.GET("/:id", orderController.GetOrder)
	}
}
