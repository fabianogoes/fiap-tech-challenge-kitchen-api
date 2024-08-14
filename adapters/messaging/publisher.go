package messaging

import (
	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"github.com/goccy/go-json"
)

type RestaurantPublisher struct {
	config       *entities.Config
	awsSqsClient *AWSSQSClient
}

func NewRestaurantPublisher(
	config *entities.Config,
	awsSqsClient *AWSSQSClient,
) *RestaurantPublisher {
	return &RestaurantPublisher{
		config,
		awsSqsClient,
	}
}

func (r *RestaurantPublisher) PublishCallback(orderID uint, status string) error {
	messageBody := toCallbackJson(orderID, status)
	return r.awsSqsClient.Publish(r.config.KitchenCallbackQueueUrl, messageBody)
}

func toCallbackJson(orderID uint, status string) string {
	jsonBytes, _ := json.Marshal(map[string]interface{}{
		"orderId": orderID,
		"status":  status,
	})

	return string(jsonBytes)
}
