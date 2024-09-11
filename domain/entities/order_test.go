package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_OrderStatusToString(t *testing.T) {
	assert.Equal(t, "KITCHEN_WAITING", OrderStatusKitchenWaiting.ToString())
	assert.Equal(t, "KITCHEN_PREPARATION", OrderStatusKitchenPreparation.ToString())
	assert.Equal(t, "KITCHEN_READY", OrderStatusKitchenReady.ToString())
	assert.Equal(t, "KITCHEN_CANCELED", OrderStatusKitchenCanceled.ToString())
	assert.Equal(t, "UNKNOWN", OrderStatusKitchenUnknown.ToString())
}

func Test_OrderStringToStatus(t *testing.T) {
	assert.Equal(t, OrderStatusKitchenWaiting, ToOrderStatus("KITCHEN_WAITING"))
	assert.Equal(t, OrderStatusKitchenPreparation, ToOrderStatus("KITCHEN_PREPARATION"))
	assert.Equal(t, OrderStatusKitchenReady, ToOrderStatus("KITCHEN_READY"))
	assert.Equal(t, OrderStatusKitchenCanceled, ToOrderStatus("KITCHEN_CANCELED"))
	assert.Equal(t, OrderStatusKitchenUnknown, ToOrderStatus("UNKNOWN"))
}
