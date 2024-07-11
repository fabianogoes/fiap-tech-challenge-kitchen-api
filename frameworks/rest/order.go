package rest

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"github.com/fabianogoes/fiap-kitchen/domain/ports"
	"github.com/fabianogoes/fiap-kitchen/frameworks/rest/dto"

	"github.com/gin-gonic/gin"
)

type KitchenHandler struct {
	UseCase ports.KitchenUseCasePort
}

func NewKitchenHandler(
	useCase ports.KitchenUseCasePort,
) *KitchenHandler {
	return &KitchenHandler{
		UseCase: useCase,
	}
}

func (h *KitchenHandler) Creation(c *gin.Context) {
	var request dto.CreationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	entity := dto.ToOrderEntity(&request)
	order, err := h.UseCase.Creation(entity)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.ToOrderResponse(order))
}

func (h *KitchenHandler) GetById(c *gin.Context) {
	log.Default().Println("GetById...")
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err := h.UseCase.GetById(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ToOrderResponse(order))
}

func (h *KitchenHandler) GetAll(c *gin.Context) {
	statusPar := c.Query("status")
	status := entities.ToOrderStatus(statusPar)
	orders, err := h.UseCase.GetAll(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if len(orders) == 0 {
		c.JSON(http.StatusNoContent, gin.H{
			"message": "No orders found",
		})
	}

	c.JSON(http.StatusOK, dto.ToOrderResponses(orders))
}

func (h *KitchenHandler) Preparation(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err := h.UseCase.Preparation(uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ToOrderResponse(order))
}

func (h *KitchenHandler) Ready(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err := h.UseCase.Ready(uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ToOrderResponse(order))
}

func (h *KitchenHandler) Cancel(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err := h.UseCase.Cancel(uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ToOrderResponse(order))
}
