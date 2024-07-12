package usecases

import (
	"fmt"
	"log"

	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"github.com/fabianogoes/fiap-kitchen/domain/ports"
)

type KitchenService struct {
	kitchenRepository ports.KitchenRepositoryPort
	restaurantClient  ports.RestaurantClientPort
}

func NewKitchenService(
	kitchenRepository ports.KitchenRepositoryPort,
	restaurantClient ports.RestaurantClientPort,
) *KitchenService {
	return &KitchenService{
		kitchenRepository: kitchenRepository,
		restaurantClient:  restaurantClient,
	}
}

func (ks *KitchenService) Creation(order *entities.Order) (*entities.Order, error) {
	fmt.Println(order)
	order.Status = entities.OrderStatusWaiting
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

	order.Status = entities.OrderStatusInPreparation
	return ks.kitchenRepository.UpdateStatus(order)
}

func (ks *KitchenService) Ready(orderID uint) (*entities.Order, error) {
	order, err := ks.GetById(orderID)
	if err != nil {
		return nil, err
	}

	order.Status = entities.OrderStatusReady
	_, err = ks.kitchenRepository.UpdateStatus(order)
	if err != nil {
		return nil, fmt.Errorf("error updating order %d status: %v", orderID, err)
	}
	fmt.Printf("order %d updated to ready successfully \n", orderID)

	err = ks.restaurantClient.ReadyForDelivery(orderID)
	if err != nil {
		return nil, fmt.Errorf("error calling restaurant ready for delivery: %v", err)
	}

	return order, nil
}

func (ks *KitchenService) Cancel(orderID uint) (*entities.Order, error) {
	order, err := ks.GetById(orderID)
	if err != nil {
		return nil, err
	}

	order.Status = entities.OrderStatusCanceled
	return ks.kitchenRepository.UpdateStatus(order)
}
