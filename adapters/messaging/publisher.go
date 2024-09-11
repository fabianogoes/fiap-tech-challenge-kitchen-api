package messaging

import (
	"fmt"
	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"github.com/fabianogoes/fiap-kitchen/frameworks/repository"
	"github.com/goccy/go-json"
)

type RestaurantPublisher struct {
	config           *entities.Config
	awsSqsClient     *AWSSQSClient
	outboxRepository *repository.OutboxRepository
}

func NewRestaurantPublisher(
	config *entities.Config,
	awsSqsClient *AWSSQSClient,
	outboxRepository *repository.OutboxRepository,
) *RestaurantPublisher {
	return &RestaurantPublisher{
		config,
		awsSqsClient,
		outboxRepository,
	}
}

func (r *RestaurantPublisher) PublishCallback(orderID uint, status string) error {
	messageBody := toCallbackJson(orderID, status)
	queueUrl := r.config.KitchenCallbackQueueUrl

	outbox, err := r.outboxRepository.CreateOutbox(orderID, messageBody, queueUrl)
	if err != nil {
		return fmt.Errorf("error creating outbox for order %v: %v", orderID, err)
	}

	err = r.awsSqsClient.Publish(queueUrl, messageBody)
	if err != nil {
		return fmt.Errorf("error sending message to kitchen: %v", err)
	}

	if err := r.outboxRepository.DeleteOutbox(outbox.ID); err != nil {
		return fmt.Errorf("error deleting outbox for order %v: %v", orderID, err)
	}

	return nil
}

func toCallbackJson(orderID uint, status string) string {
	jsonBytes, _ := json.Marshal(map[string]interface{}{
		"orderId": orderID,
		"status":  status,
	})

	return string(jsonBytes)
}
