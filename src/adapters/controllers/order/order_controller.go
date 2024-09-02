package controllers

import (
	"fmt"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/dtos"
	order "github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases/order"
	"github.com/Food-fusion-Fiap/order-service/src/infra/db/repositories"
	"github.com/Food-fusion-Fiap/order-service/src/infra/web/http-clients/customer"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	var inputDto dtos.CreateOrderDto

	if err := c.ShouldBindJSON(&inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := validator.Validate(inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	orderRepository := &repositories.OrderRepository{}
	customerRepository := &customer.CustomerClient{}
	productRepository := &repositories.ProductRepository{}

	usecase := order.CreateOrderUsecase{
		OrderRepository:    orderRepository,
		CustomerRepository: customerRepository,
		ProductRepository:  productRepository,
	}

	result, err := usecase.Execute(inputDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func ChangeOrderStatus(c *gin.Context) {
	var inputDto dtos.ChangeOrderStatusDto
	id := c.Params.ByName("id")

	if err := c.ShouldBindJSON(&inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := validator.Validate(inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	orderRepository := &repositories.OrderRepository{}

	usecase := order.ChangeOrderStatusUsecase{
		OrderRepository: orderRepository,
	}

	orderResult, err := usecase.Execute(id, inputDto.ChangeToStatus)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orderId":   orderResult.ID,
		"status":    orderResult.Status,
		"updatedAt": orderResult.UpdatedAt,
	})
}

func ListOngoingOrders(c *gin.Context) {
	orderRepository := &repositories.OrderRepository{}

	usecase := order.ListOngoingOrdersUsecase{
		OrderRepository: orderRepository,
	}

	orders, err := usecase.Execute()

	if err != nil {
		fmt.Println("there was an error to process the usecase", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if orders == nil {
		c.JSON(http.StatusOK, []string{})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	id := c.Params.ByName("id")

	orderRepository := &repositories.OrderRepository{}

	usecase := order.GetOrderUseCase{
		OrderRepository: orderRepository,
	}

	orders, err := usecase.Execute(id)

	if err != nil {
		fmt.Println("there was an error to process GetOrderUseCase", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if orders == nil {
		c.JSON(http.StatusOK, []string{})
		return
	}

	c.JSON(http.StatusOK, orders)
}
