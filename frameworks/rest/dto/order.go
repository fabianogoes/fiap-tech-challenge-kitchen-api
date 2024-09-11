package dto

import (
	"github.com/fabianogoes/fiap-kitchen/domain/entities"
)

type CreationRequest struct {
	ID    uint         `json:"id"`
	Items []*OrderItem `json:"items"`
}

type OrderItem struct {
	ID       uint     `json:"id"`
	Product  *Product `json:"product"`
	Quantity int      `json:"quantity"`
}

type Product struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type OrderResponse struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}

func ToOrderEntity(dto *CreationRequest) *entities.Order {
	items := make([]*entities.OrderItem, 0)
	for _, item := range dto.Items {
		items = append(items, &entities.OrderItem{
			Product: &entities.Product{
				Name:     item.Product.Name,
				Category: item.Product.Category,
			},
			Quantity: item.Quantity,
		})
	}
	return &entities.Order{
		ID:     dto.ID,
		Status: entities.OrderStatusKitchenWaiting,
	}
}

func ToOrderResponse(order *entities.Order) OrderResponse {
	return OrderResponse{
		ID:     order.ID,
		Status: order.Status.ToString(),
	}
}

func ToOrderResponses(orders []*entities.Order) []OrderResponse {
	var response []OrderResponse
	for _, order := range orders {
		response = append(response, ToOrderResponse(order))
	}
	return response
}
