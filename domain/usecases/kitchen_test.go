package usecases

import (
	"errors"
	"fmt"
	"testing"

	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"github.com/stretchr/testify/mock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKitchen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kitchen Suite")
}

var _ = Describe("Kitchen", func() {
	Context("creation success", func() {
		orderWaiting := *OrderWithID
		orderWaiting.Status = entities.OrderStatusKitchenWaiting
		kitchenRepositoryMock := new(KitchenRepositoryMock)
		kitchenRepositoryMock.On("Create", mock.Anything).Return(&orderWaiting, nil)
		kitchenRepositoryMock.On("GetById", mock.Anything).Return(nil, nil)

		restaurantPublisher := new(RestaurantPublisherMock)
		restaurantPublisher.On("PublishCallback", mock.Anything).Return(nil)

		useCase := NewKitchenService(kitchenRepositoryMock, restaurantPublisher)

		order, err := useCase.Creation(&orderWaiting)

		It("has no error on creation", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not be nil", func() {
			Expect(order).ShouldNot(BeNil())
		})

		It(fmt.Sprintf("has id %d", order.ID), func() {
			Expect(order.ID).Should(Equal(OrderWithID.ID))
		})
	})

	Context("creation failed", func() {
		kitchenRepositoryMock := new(KitchenRepositoryMock)
		kitchenRepositoryMock.On("Create", mock.Anything).Return(nil, errors.New("creation error"))
		kitchenRepositoryMock.On("GetById", mock.Anything).Return(nil, nil)

		restaurantPublisher := new(RestaurantPublisherMock)
		restaurantPublisher.On("PublishCallback", mock.Anything).Return(nil)

		useCase := NewKitchenService(kitchenRepositoryMock, restaurantPublisher)

		orderToCreation := *OrderWithoutID
		orderToCreation.ID = OrderIdFail
		order, err := useCase.Creation(&orderToCreation)
		It("has no error on creation", func() {
			Expect(err).ShouldNot(BeNil())
		})

		It("has order not be nil", func() {
			Expect(order).Should(BeNil())
		})
	})

	Context("get by id success", func() {
		kitchenRepositoryMock := new(KitchenRepositoryMock)
		kitchenRepositoryMock.On("GetById", mock.Anything).Return(OrderWithID, nil)

		restaurantPublisher := new(RestaurantPublisherMock)
		restaurantPublisher.On("PublishCallback", mock.Anything).Return(nil)

		useCase := NewKitchenService(kitchenRepositoryMock, restaurantPublisher)

		order, err := useCase.GetById(OrderWithID.ID)
		It("has no error on get", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not be nil", func() {
			Expect(order).ShouldNot(BeNil())
		})
	})

	Context("get by id failed", func() {
		kitchenRepositoryMock := new(KitchenRepositoryMock)
		kitchenRepositoryMock.On("GetById", mock.Anything).Return(nil, errors.New("get error"))

		restaurantPublisher := new(RestaurantPublisherMock)
		restaurantPublisher.On("PublishCallback", mock.Anything).Return(nil)

		useCase := NewKitchenService(kitchenRepositoryMock, restaurantPublisher)

		order, err := useCase.GetById(OrderIdFail)
		It("has error on get", func() {
			Expect(err).ShouldNot(BeNil())
		})

		It("has order not be nil", func() {
			Expect(order).Should(BeNil())
		})
	})

	Context("preparation success", func() {
		orderPreparation := *OrderWaiting
		orderPreparation.Status = entities.OrderStatusKitchenPreparation

		kitchenRepositoryMock := new(KitchenRepositoryMock)
		kitchenRepositoryMock.On("GetById", OrderWaiting.ID).Return(OrderWaiting, nil)
		kitchenRepositoryMock.On("UpdateStatus", OrderWaiting).Return(&orderPreparation, nil)

		restaurantPublisher := new(RestaurantPublisherMock)
		restaurantPublisher.On("PublishCallback", mock.Anything).Return(nil)

		useCase := NewKitchenService(kitchenRepositoryMock, restaurantPublisher)

		order, err := useCase.Preparation(orderPreparation.ID)
		It("has no error on preparation", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not nil", func() {
			Expect(order).ShouldNot(BeNil())
		})

		It(fmt.Sprintf("has order with status %v", entities.OrderStatusKitchenPreparation), func() {
			Expect(order.Status).Should(Equal(entities.OrderStatusKitchenPreparation))
		})
	})

	Context("preparation failed", func() {
		kitchenRepositoryMock := new(KitchenRepositoryMock)
		kitchenRepositoryMock.On("GetById", OrderIdFail).Return(nil, errors.New("not found"))

		restaurantPublisher := new(RestaurantPublisherMock)
		restaurantPublisher.On("PublishCallback", mock.Anything).Return(nil)

		useCase := NewKitchenService(kitchenRepositoryMock, restaurantPublisher)

		orderFail, err := useCase.Preparation(OrderIdFail)
		It("has error on preparation", func() {
			Expect(err).ShouldNot(BeNil())
		})

		It("has order be nil", func() {
			Expect(orderFail).Should(BeNil())
		})

	})

	Context("ready success", func() {
		kitchenRepositoryMock := new(KitchenRepositoryMock)

		kitchenRepositoryMock.On("GetById", OrderInPreparation.ID).Return(OrderInPreparation, nil)
		kitchenRepositoryMock.On("UpdateStatus", mock.Anything).Return(OrderReady, nil)

		restaurantPublisher := new(RestaurantPublisherMock)
		restaurantPublisher.On("PublishCallback", mock.Anything).Return(nil)

		useCase := NewKitchenService(kitchenRepositoryMock, restaurantPublisher)

		inReady, err := useCase.Ready(OrderInPreparation.ID)
		It("has no error on ready", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not nil", func() {
			Expect(inReady).ShouldNot(BeNil())
		})

	})

	Context("ready failed", func() {
		kitchenRepositoryMock := new(KitchenRepositoryMock)
		kitchenRepositoryMock.On("GetById", OrderIdFail).Return(nil, errors.New("not found"))

		restaurantPublisher := new(RestaurantPublisherMock)
		restaurantPublisher.On("PublishCallback", mock.Anything).Return(nil)

		useCase := NewKitchenService(kitchenRepositoryMock, restaurantPublisher)

		orderFail, err := useCase.Ready(OrderIdFail)
		It("has error on ready", func() {
			Expect(err).ShouldNot(BeNil())
		})

		It("has order nil", func() {
			Expect(orderFail).Should(BeNil())
		})
	})

	Context("cancel failed", func() {
		kitchenRepositoryMock := new(KitchenRepositoryMock)
		kitchenRepositoryMock.On("GetById", OrderIdFail).Return(nil, errors.New("not found"))

		restaurantPublisher := new(RestaurantPublisherMock)
		restaurantPublisher.On("PublishCallback", mock.Anything).Return(nil)

		useCase := NewKitchenService(kitchenRepositoryMock, restaurantPublisher)

		orderFail, err := useCase.Cancel(OrderIdFail)
		It("has error on cancel", func() {
			Expect(err).ShouldNot(BeNil())
		})

		It("has order nil", func() {
			Expect(orderFail).Should(BeNil())
		})
	})

	Context("cancel success", func() {
		kitchenRepositoryMock := new(KitchenRepositoryMock)

		kitchenRepositoryMock.On("GetById", OrderInPreparation.ID).Return(OrderInPreparation, nil)
		kitchenRepositoryMock.On("UpdateStatus", OrderCanceled).Return(OrderCanceled, nil)

		restaurantPublisher := new(RestaurantPublisherMock)
		restaurantPublisher.On("PublishCallback", mock.Anything).Return(nil)

		useCase := NewKitchenService(kitchenRepositoryMock, restaurantPublisher)

		order, err := useCase.Cancel(OrderInPreparation.ID)
		It("has no error on cancel", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not nil", func() {
			Expect(order).ShouldNot(BeNil())
		})
	})

	Context("update failed", func() {
		order := OrderWithoutID
		order.ID = OrderIdFail
		kitchenRepositoryMock := new(KitchenRepositoryMock)

		kitchenRepositoryMock.On("GetById", mock.Anything).Return(nil, errors.New("not found"))
		kitchenRepositoryMock.On("UpdateStatus", order).Return(nil, errors.New("update error"))

		restaurantPublisher := new(RestaurantPublisherMock)
		restaurantPublisher.On("PublishCallback", mock.Anything).Return(nil)

		useCase := NewKitchenService(kitchenRepositoryMock, restaurantPublisher)

		order, err := useCase.Ready(OrderIdFail)
		It("has error on update", func() {
			Expect(err).ShouldNot(BeNil())
		})

		It("has order not nil", func() {
			Expect(order).Should(BeNil())
		})
	})
})
