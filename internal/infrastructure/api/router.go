package api

import (
	"notifier/internal/infrastructure/api/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	setupSwagger(router)
	setupRoutes(router)

	return router
}

func setupSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func setupRoutes(router *gin.Engine) {
	router.GET("/subscriptions/:id", handlers.GetSubscription)
	router.POST("/subscriptions", handlers.CreateSubscription)
}
