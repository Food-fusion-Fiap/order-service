package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()

	SetupProductRoutes(router)
	SetupOrderRoutes(router)

	err := router.Run()

	if err != nil {
		log.Panic(err)
		return
	}
}
