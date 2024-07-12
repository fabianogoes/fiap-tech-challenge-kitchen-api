package ports

type RestaurantClientPort interface {
	ReadyForDelivery(orderID uint) error
}
