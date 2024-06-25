package usecases

import (
	"fmt"
	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"github.com/fabianogoes/fiap-kitchen/domain/ports"
)

type KitchenService struct {
	kitchenRepository ports.KitchenRepositoryPort
}

func NewKitchenService(
	kitchenRepository ports.KitchenRepositoryPort,
) *KitchenService {
	return &KitchenService{
		kitchenRepository: kitchenRepository,
	}
}

func (ks *KitchenService) Creation(order *entities.Order) (*entities.Order, error) {
	fmt.Println(order)
	return ks.kitchenRepository.Create(order)
}

func (ks *KitchenService) GetById(id uint) (*entities.Order, error) {
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
	return ks.kitchenRepository.UpdateStatus(order)
}

func (ks *KitchenService) Cancel(orderID uint) (*entities.Order, error) {
	order, err := ks.GetById(orderID)
	if err != nil {
		return nil, err
	}

	order.Status = entities.OrderStatusCanceled
	return ks.kitchenRepository.UpdateStatus(order)
}
