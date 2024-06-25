package rest

import (
	"github.com/fabianogoes/fiap-kitchen/domain/usecases"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Router(t *testing.T) {
	kitchenService := usecases.NewKitchenService(new(usecases.KitchenRepositoryMock))
	kitchenHandler := NewKitchenHandler(kitchenService)
	router, err := NewRouter(kitchenHandler)
	assert.Nil(t, err)
	assert.NotNil(t, router)
}
