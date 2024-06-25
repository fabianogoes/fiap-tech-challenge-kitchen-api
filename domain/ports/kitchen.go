package ports

import "github.com/fabianogoes/fiap-kitchen/domain/entities"

type KitchenUseCasePort interface {
	Creation(order *entities.Order) (*entities.Order, error)
	GetById(id uint) (*entities.Order, error)
	GetAll(status entities.OrderStatus) ([]*entities.Order, error)
	Preparation(orderID uint) (*entities.Order, error)
	Ready(orderID uint) (*entities.Order, error)
	Cancel(orderID uint) (*entities.Order, error)
}

type KitchenRepositoryPort interface {
	Create(order *entities.Order) (*entities.Order, error)
	GetById(id uint) (*entities.Order, error)
	GetAll(status entities.OrderStatus) ([]*entities.Order, error)
	UpdateStatus(order *entities.Order) (*entities.Order, error)
}
