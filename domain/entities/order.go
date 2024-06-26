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
	OrderStatusWaiting OrderStatus = iota
	OrderStatusInPreparation
	OrderStatusReady
	OrderStatusCanceled
	OrderStatusUnknown
)

func (os OrderStatus) ToString() string {
	return [...]string{
		"WAITING",
		"PREPARATION",
		"READY",
		"CANCELED",
		"UNKNOWN",
	}[os]
}

func ToOrderStatus(status string) OrderStatus {
	switch status {
	case "WAITING":

		return OrderStatusWaiting
	case "PREPARATION":

		return OrderStatusInPreparation
	case "READY":

		return OrderStatusReady
	case "CANCELED":

		return OrderStatusCanceled
	default:

		return OrderStatusUnknown
	}
}
