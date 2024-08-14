package usecases

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"github.com/fabianogoes/fiap-kitchen/domain/ports"
)

type KitchenService struct {
	kitchenRepository   ports.KitchenRepositoryPort
	restaurantPublisher ports.RestaurantPublisherPort
}

func NewKitchenService(
	kitchenRepository ports.KitchenRepositoryPort,
	restaurantPublisher ports.RestaurantPublisherPort,
) *KitchenService {
	return &KitchenService{
		kitchenRepository:   kitchenRepository,
		restaurantPublisher: restaurantPublisher,
	}
}

func (ks *KitchenService) Creation(order *entities.Order) (*entities.Order, error) {
	idempotency, err := ks.GetById(order.ID)
	if err == nil && idempotency != nil {
		slog.Warn(fmt.Sprintf("idempotency validation, kitchen with orderId %v already exists", order.ID))
		return idempotency, nil
	}

	order.Status = entities.OrderStatusKitchenWaiting
	return ks.kitchenRepository.Create(order)
}

func (ks *KitchenService) GetById(id uint) (*entities.Order, error) {
	log.Default().Printf("GetById orderID: %d \n", id)
	return ks.kitchenRepository.GetById(id)
}

func (ks *KitchenService) GetAll(status entities.OrderStatus) ([]*entities.Order, error) {
	return ks.kitchenRepository.GetAll(status)
}

func (ks *KitchenService) Preparation(orderID uint) (*entities.Order, error) {
	order, err := ks.GetById(orderID)
	if err != nil {
		return nil, err
	}

	order.Status = entities.OrderStatusKitchenPreparation
	orderPreparation, err := ks.kitchenRepository.UpdateStatus(order)
	if err != nil {
		return nil, fmt.Errorf("error updating order %d status: %v - %v\n", orderID, order.Status.ToString(), err)
	}

	err = ks.restaurantPublisher.PublishCallback(orderID, order.Status.ToString())
	if err != nil {
		return nil, fmt.Errorf("error calling restaurant ready for delivery: %v\n", err)
	}

	return orderPreparation, nil
}

func (ks *KitchenService) Ready(orderID uint) (*entities.Order, error) {
	order, err := ks.GetById(orderID)
	if err != nil {
		return nil, err
	}

	order.Status = entities.OrderStatusKitchenReady
	_, err = ks.kitchenRepository.UpdateStatus(order)
	if err != nil {
		return nil, fmt.Errorf("error updating order %d status: %v - %v\n", orderID, order.Status.ToString(), err)
	}
	fmt.Printf("order %d updated to ready successfully \n", orderID)

	err = ks.restaurantPublisher.PublishCallback(orderID, order.Status.ToString())
	if err != nil {
		return nil, fmt.Errorf("error calling restaurant ready for delivery: %v\n", err)
	}

	return order, nil
}

func (ks *KitchenService) Cancel(orderID uint) (*entities.Order, error) {
	order, err := ks.GetById(orderID)
	if err != nil {
		return nil, err
	}

	order.Status = entities.OrderStatusKitchenCanceled
	return ks.kitchenRepository.UpdateStatus(order)
}
