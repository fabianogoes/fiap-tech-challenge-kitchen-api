package dbo

import (
	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"time"
)

type Order struct {
	ID        uint         `bson:"orderId"`
	CreatedAt time.Time    `bson:"created_at"`
	UpdatedAt time.Time    `bson:"updated_at"`
	Status    string       `bson:"status"`
	Items     []*OrderItem `bson:"items"`
}

func (o *Order) ToOrderEntity() *entities.Order {
	var items []*entities.OrderItem

	for _, item := range o.Items {
		items = append(items, item.ToItemEntity())
	}

	return &entities.Order{
		ID:     o.ID,
		Status: o.ToOrderStatus(),
		Items:  items,
	}
}

func (o *Order) ToOrderStatus() entities.OrderStatus {
	switch o.Status {
	case "KITCHEN_WAITING":
		return entities.OrderStatusKitchenWaiting
	case "KITCHEN_PREPARATION":
		return entities.OrderStatusKitchenPreparation
	case "KITCHEN_READY":
		return entities.OrderStatusKitchenReady
	case "KITCHEN_CANCELED":
		return entities.OrderStatusKitchenCanceled
	default:
		return entities.OrderStatusKitchenUnknown
	}
}

type OrderItem struct {
	Product  *Product `bson:"product"`
	Quantity int      `bson:"quantity"`
}

type Product struct {
	Name     string `bson:"name"`
	Category string `bson:"category"`
}

func (i *OrderItem) ToItemEntity() *entities.OrderItem {
	return &entities.OrderItem{
		Product: &entities.Product{
			Name:     i.Product.Name,
			Category: i.Product.Category,
		},
		Quantity: int(i.Quantity),
	}
}

func ToOrderDBO(order *entities.Order) *Order {
	items := make([]*OrderItem, len(order.Items))
	for _, item := range order.Items {
		items = append(items, &OrderItem{
			Product: &Product{
				Name:     item.Product.Name,
				Category: item.Product.Category,
			},
			Quantity: item.Quantity,
		})
	}

	return &Order{
		ID:        order.ID,
		CreatedAt: time.Now(),
		Status:    order.Status.ToString(),
		Items:     items,
	}
}
