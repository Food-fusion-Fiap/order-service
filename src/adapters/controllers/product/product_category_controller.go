package controllers

import (
	usecasesProduct "github.com/Food-fusion-Fiap/order-service/src/core/domain/usecases/product"
	"github.com/Food-fusion-Fiap/order-service/src/infra/db/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ListCategory(ctx *gin.Context) {
	productRepository := repositories.ProductCategoryRepository{}

	result, err := usecasesProduct.BuildListProductCategoryUsecase(productRepository).Execute()

	if err != nil {
		log.Println("there was an error to retrieve products", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
