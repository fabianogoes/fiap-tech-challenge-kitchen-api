package usecases

import (
	"errors"
	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestKitchenService_Creation(t *testing.T) {
	kitchenRepositoryMock := new(KitchenRepositoryMock)
	kitchenRepositoryMock.On("Create", mock.Anything).Return(OrderWithID)

	useCase := NewKitchenService(kitchenRepositoryMock)

	creation, err := useCase.Creation(OrderWithoutID)
	assert.NoError(t, err)
	assert.NotNil(t, creation)
	assert.NotNil(t, creation.ID)
	assert.Equal(t, entities.OrderStatusWaiting, creation.Status)
}

func TestKitchenService_CreationFail(t *testing.T) {
	order := OrderWithoutID
	order.ID = OrderIdFail

	kitchenRepositoryMock := new(KitchenRepositoryMock)
	kitchenRepositoryMock.On("Create", order).Return(nil, errors.New("creation error"))

	useCase := NewKitchenService(kitchenRepositoryMock)

	creation, err := useCase.Creation(order)
	assert.Error(t, err)
	assert.Nil(t, creation)
}

func TestKitchenService_GetById(t *testing.T) {
	order := OrderWithoutID
	order.ID = OrderIdFail

	kitchenRepositoryMock := new(KitchenRepositoryMock)
	kitchenRepositoryMock.On("GetById", order.ID).Return(nil, errors.New("not found"))

	useCase := NewKitchenService(kitchenRepositoryMock)

	order, err := useCase.GetById(order.ID)
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestKitchenService_GetByIdFail(t *testing.T) {
	kitchenRepositoryMock := new(KitchenRepositoryMock)
	kitchenRepositoryMock.On("GetById", mock.Anything).Return(OrderWithID)

	useCase := NewKitchenService(kitchenRepositoryMock)

	order, err := useCase.GetById(OrderIdSuccess)
	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.NotNil(t, order.ID)
	assert.Equal(t, entities.OrderStatusWaiting, order.Status)
}

func TestKitchenService_GetAll(t *testing.T) {
	kitchenRepositoryMock := new(KitchenRepositoryMock)
	list := []*entities.Order{OrderWithID}
	kitchenRepositoryMock.On("GetAll", mock.Anything).Return(list)

	useCase := NewKitchenService(kitchenRepositoryMock)

	orders, err := useCase.GetAll(entities.OrderStatusReady)
	assert.NoError(t, err)
	assert.NotNil(t, orders)
	assert.Len(t, orders, 1)
}

func TestKitchenService_Preparation(t *testing.T) {
	kitchenRepositoryMock := new(KitchenRepositoryMock)

	kitchenRepositoryMock.On("GetById", OrderWaiting.ID).Return(OrderWaiting, nil)
	kitchenRepositoryMock.On("UpdateStatus", OrderWaiting).Return(OrderInPreparation, nil)

	useCase := NewKitchenService(kitchenRepositoryMock)

	inPreparation, err := useCase.Preparation(OrderWaiting.ID)
	assert.NoError(t, err)
	assert.NotNil(t, inPreparation)
	assert.NotNil(t, inPreparation.ID)
	assert.Equal(t, entities.OrderStatusInPreparation, inPreparation.Status)
}

func TestKitchenService_PreparationFail(t *testing.T) {
	kitchenRepositoryMock := new(KitchenRepositoryMock)
	kitchenRepositoryMock.On("GetById", OrderIdFail).Return(nil, errors.New("not found"))

	useCase := NewKitchenService(kitchenRepositoryMock)

	orderFail, err := useCase.Preparation(OrderIdFail)
	assert.Error(t, err)
	assert.Nil(t, orderFail)
}

func TestKitchenService_Ready(t *testing.T) {
	kitchenRepositoryMock := new(KitchenRepositoryMock)

	kitchenRepositoryMock.On("GetById", OrderInPreparation.ID).Return(OrderInPreparation, nil)
	kitchenRepositoryMock.On("UpdateStatus", mock.Anything).Return(OrderReady, nil)

	useCase := NewKitchenService(kitchenRepositoryMock)

	inReady, err := useCase.Ready(OrderInPreparation.ID)
	assert.NoError(t, err)
	assert.NotNil(t, inReady)
	assert.NotNil(t, inReady.ID)
	assert.Equal(t, entities.OrderStatusReady, inReady.Status)
}

func TestKitchenService_ReadyFail(t *testing.T) {
	kitchenRepositoryMock := new(KitchenRepositoryMock)
	kitchenRepositoryMock.On("GetById", OrderIdFail).Return(nil, errors.New("not found"))

	useCase := NewKitchenService(kitchenRepositoryMock)

	orderFail, err := useCase.Ready(OrderIdFail)
	assert.Error(t, err)
	assert.Nil(t, orderFail)
}

func TestKitchenService_CancelFail(t *testing.T) {
	kitchenRepositoryMock := new(KitchenRepositoryMock)
	kitchenRepositoryMock.On("GetById", OrderIdFail).Return(nil, errors.New("not found"))

	useCase := NewKitchenService(kitchenRepositoryMock)

	orderFail, err := useCase.Cancel(OrderIdFail)
	assert.Error(t, err)
	assert.Nil(t, orderFail)
}

func TestKitchenService_Cancel(t *testing.T) {
	kitchenRepositoryMock := new(KitchenRepositoryMock)

	kitchenRepositoryMock.On("GetById", OrderInPreparation.ID).Return(OrderInPreparation, nil)
	kitchenRepositoryMock.On("UpdateStatus", OrderCanceled).Return(OrderCanceled, nil)

	useCase := NewKitchenService(kitchenRepositoryMock)

	inReady, err := useCase.Cancel(OrderInPreparation.ID)
	assert.NoError(t, err)
	assert.NotNil(t, inReady)
	assert.NotNil(t, inReady.ID)
	assert.Equal(t, entities.OrderStatusCanceled, inReady.Status)
}

func TestKitchenService_UpdateFail(t *testing.T) {
	order := OrderWithoutID
	order.ID = OrderIdFail
	kitchenRepositoryMock := new(KitchenRepositoryMock)

	kitchenRepositoryMock.On("GetById", mock.Anything).Return(nil, errors.New("not found"))
	kitchenRepositoryMock.On("UpdateStatus", order).Return(nil, errors.New("update error"))

	useCase := NewKitchenService(kitchenRepositoryMock)

	inReady, err := useCase.Ready(OrderIdFail)
	assert.Error(t, err)
	assert.Nil(t, inReady)
}
