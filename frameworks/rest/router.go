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

	orders := router.Group("/orders")
	{
		orders.POST("/", kitchenHandler.Creation)
		orders.GET("/:id", kitchenHandler.GetById)
		orders.GET("/", kitchenHandler.GetAll)
		orders.PUT("/:id/preparation", kitchenHandler.Preparation)
		orders.PUT("/:id/ready", kitchenHandler.Ready)
		orders.PUT("/:id/cancel", kitchenHandler.Cancel)
	}

	return &Router{
		router,
	}, nil
}
