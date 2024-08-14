package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"github.com/fabianogoes/fiap-kitchen/domain/usecases"
	"github.com/fabianogoes/fiap-kitchen/frameworks/rest/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var creationRequestSuccess = dto.CreationRequest{
	ID: 1,
	Items: []*dto.OrderItem{
		{
			ID: usecases.OrderIdSuccess,
			Product: &dto.Product{
				ID:       1,
				Name:     "Test",
				Category: "Test",
			},
		},
	},
}

var creationRequestFail = dto.CreationRequest{
	ID: 1,
	Items: []*dto.OrderItem{
		{
			ID: usecases.OrderIdFail,
			Product: &dto.Product{
				ID:       1,
				Name:     "Test",
				Category: "Test",
			},
		},
	},
}

func TestKitchenHandler_Creation(t *testing.T) {
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("Create", mock.Anything).Return(usecases.OrderWithID)

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	jsonRequest, _ := json.Marshal(creationRequestSuccess)
	readerRequest := bytes.NewReader(jsonRequest)

	r := SetupTest()
	r.POST("/kitchen/orders/", handler.Creation)
	request, _ := http.NewRequest("POST", "/kitchen/orders/", readerRequest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 201, response.Code, "CREATE response status is expected")
}

func TestKitchenHandler_CreationFailBdRequest(t *testing.T) {
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("Create", mock.Anything).Return(nil, errors.New("creation error"))

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	r := SetupTest()
	r.POST("/kitchen/orders/", handler.Creation)
	request, _ := http.NewRequest("POST", "/kitchen/orders/", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 400, response.Code)
}

func TestKitchenHandler_CreationFail(t *testing.T) {
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("Create", mock.Anything).Return(nil, errors.New("creation error"))

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	creationRequestFail.ID = usecases.OrderIdFail
	jsonRequest, _ := json.Marshal(creationRequestFail)
	readerRequest := bytes.NewReader(jsonRequest)

	r := SetupTest()
	r.POST("/kitchen/orders/", handler.Creation)
	request, _ := http.NewRequest("POST", "/kitchen/orders/", readerRequest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 500, response.Code, "CREATE response status is expected")
}

func TestKitchenHandler_GetById(t *testing.T) {
	orderReady := usecases.OrderReady
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("GetById", mock.Anything).Return(orderReady, nil)

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	r := SetupTest()
	r.GET("/kitchen/orders/:id", handler.GetById)
	request, _ := http.NewRequest("GET", fmt.Sprintf("/kitchen/orders/%d", orderReady.ID), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}

func TestKitchenHandler_GetByIdFailNotFound(t *testing.T) {
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("GetById", mock.Anything).Return(nil, errors.New("not found"))

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	r := SetupTest()
	r.GET("/kitchen/orders/:id", handler.GetById)
	request, _ := http.NewRequest("GET", fmt.Sprintf("/kitchen/orders/%d", usecases.OrderIdFail), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 404, response.Code)
}

func TestKitchenHandler_GetByIdFailBadRequest(t *testing.T) {
	repository := new(usecases.KitchenRepositoryMock)

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	r := SetupTest()
	r.GET("/kitchen/orders/:id", handler.GetById)
	request, _ := http.NewRequest("GET", "/kitchen/orders/a", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 400, response.Code)
}

func TestKitchenHandler_GetAll(t *testing.T) {
	orders := []*entities.Order{usecases.OrderReady}
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("GetAll", entities.OrderStatusKitchenReady).Return(orders, nil)

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	statusRequest := entities.OrderStatusKitchenReady.ToString()
	r := SetupTest()
	r.GET("/kitchen/orders/", handler.GetAll)
	request, _ := http.NewRequest("GET", fmt.Sprintf("/kitchen/orders/?status=%s", statusRequest), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}

func TestKitchenHandler_GetAllFail(t *testing.T) {
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("GetAll", entities.OrderStatusKitchenUnknown).Return(nil, errors.New("not found"))

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	statusRequest := entities.OrderStatusKitchenUnknown.ToString()
	r := SetupTest()
	r.GET("/kitchen/orders/", handler.GetAll)
	request, _ := http.NewRequest("GET", fmt.Sprintf("/kitchen/orders/?status=%s", statusRequest), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 500, response.Code)
}

func TestKitchenHandler_Preparation(t *testing.T) {
	order := usecases.OrderWaiting
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("GetById", order.ID).Return(order, nil)
	repository.On("UpdateStatus", order).Return(order, nil)

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	r := SetupTest()
	r.POST("/kitchen/orders/:id/preparation", handler.Preparation)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/kitchen/orders/%d/preparation", order.ID), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}

func TestKitchenHandler_PreparationFailBadRequest(t *testing.T) {
	order := usecases.OrderWaiting
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("GetById", order.ID).Return(order, nil)
	repository.On("UpdateStatus", order).Return(order, nil)

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	r := SetupTest()
	r.POST("/kitchen/orders/:id/preparation", handler.Preparation)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/kitchen/orders/%s/preparation", "a"), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 400, response.Code)
}

func TestKitchenHandler_PreparationFailError(t *testing.T) {
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("GetById", usecases.OrderIdFail).Return(nil, errors.New("not found"))

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	r := SetupTest()
	r.POST("/kitchen/orders/:id/preparation", handler.Preparation)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/kitchen/orders/%d/preparation", usecases.OrderIdFail), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 500, response.Code)
}

func TestKitchenHandler_Ready(t *testing.T) {
	order := usecases.OrderReady
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("GetById", order.ID).Return(order, nil)
	repository.On("UpdateStatus", order).Return(order, nil)

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	r := SetupTest()
	r.POST("/kitchen/orders/:id/ready", handler.Ready)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/kitchen/orders/%d/ready", order.ID), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}

func TestKitchenHandler_ReadyFailBadRequest(t *testing.T) {
	order := usecases.OrderReady
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("GetById", order.ID).Return(order, nil)
	repository.On("UpdateStatus", order).Return(order, nil)

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	r := SetupTest()
	r.POST("/kitchen/orders/:id/ready", handler.Ready)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/kitchen/orders/%s/ready", "a"), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 400, response.Code)
}

func TestKitchenHandler_ReadyFailError(t *testing.T) {
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("GetById", usecases.OrderIdFail).Return(nil, errors.New("not found"))

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	r := SetupTest()
	r.POST("/kitchen/orders/:id/ready", handler.Ready)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/kitchen/orders/%d/ready", usecases.OrderIdFail), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 500, response.Code)
}

func TestKitchenHandler_Cancel(t *testing.T) {
	order := usecases.OrderCanceled
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("GetById", order.ID).Return(order, nil)
	repository.On("UpdateStatus", order).Return(order, nil)

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	r := SetupTest()
	r.POST("/kitchen/orders/:id/cancel", handler.Cancel)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/kitchen/orders/%d/cancel", order.ID), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}

func TestKitchenHandler_CancelFailBadRequest(t *testing.T) {
	order := usecases.OrderCanceled
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("GetById", order.ID).Return(order, nil)
	repository.On("UpdateStatus", order).Return(order, nil)

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	r := SetupTest()
	r.POST("/kitchen/orders/:id/cancel", handler.Cancel)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/kitchen/orders/%s/cancel", "a"), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 400, response.Code)
}

func TestKitchenHandler_CancelFailError(t *testing.T) {
	repository := new(usecases.KitchenRepositoryMock)
	repository.On("GetById", usecases.OrderIdFail).Return(nil, errors.New("not found"))

	useCase := usecases.NewKitchenService(repository, new(usecases.RestaurantClientMock))
	handler := NewKitchenHandler(useCase)

	r := SetupTest()
	r.POST("/kitchen/orders/:id/cancel", handler.Cancel)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/kitchen/orders/%d/cancel", usecases.OrderIdFail), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 500, response.Code)
}
