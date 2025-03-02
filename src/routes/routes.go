package routes

import (
	"bank-ledger/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Account routes
	router.POST("/accounts", controllers.CreateAccount)
	router.GET("/accounts/:id", controllers.GetAccount)

	// Transaction routes
	router.POST("/transactions", controllers.ProcessTransaction)
}
