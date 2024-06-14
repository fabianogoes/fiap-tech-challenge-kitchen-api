package rest

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	kitchenHandler *KitchenHandler,
) (*Router, error) {
	router := gin.Default()

	router.GET("/", Welcome)
	router.GET("/health", Health)
	router.GET("/env", Environment)

	orders := router.Group("/kitchen/orders")
	{
		orders.POST("/", kitchenHandler.Creation)
		orders.GET("/:id", kitchenHandler.GetById)
		orders.GET("/", kitchenHandler.GetAll)
		orders.POST("/:id/preparation", kitchenHandler.Preparation)
		orders.POST("/:id/ready", kitchenHandler.Ready)
		orders.POST("/:id/cancel", kitchenHandler.Cancel)
	}

	return &Router{
		router,
	}, nil
}
