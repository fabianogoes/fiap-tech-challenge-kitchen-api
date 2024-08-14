package usecases

import (
	"fmt"

	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"github.com/stretchr/testify/mock"
)

type KitchenRepositoryMock struct {
	mock.Mock
}

var OrderIdSuccess uint = 1
var OrderIdFail uint = 2
var OrderWithoutID = &entities.Order{
	Status: entities.OrderStatusKitchenWaiting,
	Items: []*entities.OrderItem{
		{
			Product: &entities.Product{
				Name:     "Product XXX",
				Category: "YYY",
			},
			Quantity: 1,
		},
	},
}

var OrderWithID = &entities.Order{
	ID:     OrderIdSuccess,
	Status: entities.OrderStatusKitchenWaiting,
	Items: []*entities.OrderItem{
		{
			Product: &entities.Product{
				Name:     "Product XXX",
				Category: "YYY",
			},
			Quantity: 1,
		},
	},
}

var OrderWaiting = OrderWithID

var OrderInPreparation = &entities.Order{
	ID:     OrderIdSuccess,
	Status: entities.OrderStatusKitchenPreparation,
	Items: []*entities.OrderItem{
		{
			Product: &entities.Product{
				Name:     "Product XXX",
				Category: "YYY",
			},
			Quantity: 1,
		},
	},
}

var OrderReady = &entities.Order{
	ID:     OrderIdSuccess,
	Status: entities.OrderStatusKitchenReady,
	Items: []*entities.OrderItem{
		{
			Product: &entities.Product{
				Name:     "Product XXX",
				Category: "YYY",
			},
			Quantity: 1,
		},
	},
}

var OrderCanceled = &entities.Order{
	ID:     OrderIdSuccess,
	Status: entities.OrderStatusKitchenCanceled,
	Items: []*entities.OrderItem{
		{
			Product: &entities.Product{
				Name:     "Product XXX",
				Category: "YYY",
			},
			Quantity: 1,
		},
	},
}

func (or *KitchenRepositoryMock) Create(order *entities.Order) (*entities.Order, error) {
	args := or.Called(order)

	if order.ID == OrderIdFail {
		return nil, args.Error(1)
	}

	return OrderWithID, nil
}

func (or *KitchenRepositoryMock) GetById(id uint) (*entities.Order, error) {
	args := or.Called(id)
	if id == OrderIdFail {
		return nil, args.Error(1)
	}

	return OrderWaiting, nil
}

func (or *KitchenRepositoryMock) GetAll(status entities.OrderStatus) ([]*entities.Order, error) {
	args := or.Called(status)

	if status == entities.OrderStatusKitchenUnknown {
		return nil, args.Error(1)
	}

	return []*entities.Order{OrderWithID}, nil
}

func (or *KitchenRepositoryMock) UpdateStatus(order *entities.Order) (*entities.Order, error) {
	args := or.Called(order)

	if order.Status == entities.OrderStatusKitchenPreparation {
		return OrderInPreparation, nil
	}

	if order.Status == entities.OrderStatusKitchenReady {
		return OrderReady, nil
	}

	if order.Status == entities.OrderStatusKitchenCanceled {
		return OrderCanceled, nil
	}

	return nil, args.Error(1)
}

type RestaurantClientMock struct {
	mock.Mock
}

func (p *RestaurantClientMock) ReadyForDelivery(orderID uint) error {
	fmt.Printf("ReadyForDelivery orderID: %d \n", orderID)
	return nil
}
