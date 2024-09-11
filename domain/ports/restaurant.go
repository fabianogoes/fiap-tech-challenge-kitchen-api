package ports

type RestaurantClientPort interface {
	ReadyForDelivery(orderID uint) error
}

type RestaurantReceiverPort interface {
	ReceiveOrder()
}

type RestaurantPublisherPort interface {
	PublishCallback(orderID uint, status string) error
}
