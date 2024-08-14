package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_OrderStatusToString(t *testing.T) {
	assert.Equal(t, "WAITING", OrderStatusKitchenWaiting.ToString())
	assert.Equal(t, "PREPARATION", OrderStatusKitchenPreparation.ToString())
	assert.Equal(t, "READY", OrderStatusKitchenReady.ToString())
	assert.Equal(t, "CANCELED", OrderStatusKitchenCanceled.ToString())
	assert.Equal(t, "UNKNOWN", OrderStatusKitchenUnknown.ToString())
}

func Test_OrderStringToStatus(t *testing.T) {
	assert.Equal(t, OrderStatusKitchenWaiting, ToOrderStatus("WAITING"))
	assert.Equal(t, OrderStatusKitchenPreparation, ToOrderStatus("PREPARATION"))
	assert.Equal(t, OrderStatusKitchenReady, ToOrderStatus("READY"))
	assert.Equal(t, OrderStatusKitchenCanceled, ToOrderStatus("CANCELED"))
	assert.Equal(t, OrderStatusKitchenUnknown, ToOrderStatus("UNKNOWN"))
}
