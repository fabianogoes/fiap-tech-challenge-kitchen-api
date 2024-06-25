package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_OrderStatusToString(t *testing.T) {
	assert.Equal(t, "WAITING", OrderStatusWaiting.ToString())
	assert.Equal(t, "PREPARATION", OrderStatusInPreparation.ToString())
	assert.Equal(t, "READY", OrderStatusReady.ToString())
	assert.Equal(t, "CANCELED", OrderStatusCanceled.ToString())
	assert.Equal(t, "UNKNOWN", OrderStatusUnknown.ToString())
}

func Test_OrderStringToStatus(t *testing.T) {
	assert.Equal(t, OrderStatusWaiting, ToOrderStatus("WAITING"))
	assert.Equal(t, OrderStatusInPreparation, ToOrderStatus("PREPARATION"))
	assert.Equal(t, OrderStatusReady, ToOrderStatus("READY"))
	assert.Equal(t, OrderStatusCanceled, ToOrderStatus("CANCELED"))
	assert.Equal(t, OrderStatusUnknown, ToOrderStatus("UNKNOWN"))
}
