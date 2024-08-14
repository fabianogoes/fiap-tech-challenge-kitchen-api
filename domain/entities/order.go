package entities

type Order struct {
	ID     uint
	Status OrderStatus
	Items  []*OrderItem
}

type OrderItem struct {
	Product  *Product
	Quantity int
}

type Product struct {
	Name     string
	Category string
}

type OrderStatus int

const (
	OrderStatusKitchenWaiting OrderStatus = iota
	OrderStatusKitchenPreparation
	OrderStatusKitchenReady
	OrderStatusKitchenCanceled
	OrderStatusKitchenUnknown
)

func (os OrderStatus) ToString() string {
	return [...]string{
		"KITCHEN_WAITING",
		"KITCHEN_PREPARATION",
		"KITCHEN_READY",
		"KITCHEN_CANCELED",
		"UNKNOWN",
	}[os]
}

func ToOrderStatus(status string) OrderStatus {
	switch status {
	case "KITCHEN_WAITING":

		return OrderStatusKitchenWaiting
	case "KITCHEN_PREPARATION":

		return OrderStatusKitchenPreparation
	case "KITCHEN_READY":

		return OrderStatusKitchenReady
	case "KITCHEN_CANCELED":

		return OrderStatusKitchenCanceled
	default:

		return OrderStatusKitchenUnknown
	}
}
