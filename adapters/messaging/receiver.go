package messaging

import (
	"fmt"
	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"github.com/fabianogoes/fiap-kitchen/domain/ports"
	"github.com/fabianogoes/fiap-kitchen/frameworks/rest/dto"
	"github.com/goccy/go-json"
	"log/slog"
)

type RestaurantReceiver struct {
	kitchenUseCase ports.KitchenUseCasePort
	config         *entities.Config
	awsSqsClient   *AWSSQSClient
}

func NewRestaurantReceiver(
	kitchenUseCase ports.KitchenUseCasePort,
	config *entities.Config,
	awsSqsClient *AWSSQSClient,
) *RestaurantReceiver {
	return &RestaurantReceiver{
		kitchenUseCase,
		config,
		awsSqsClient,
	}
}

func (r *RestaurantReceiver) ReceiveOrder() {
	slog.Info("running kitchen messaging listener...")
	if messages := r.awsSqsClient.Receive(r.config.KitchenQueueUrl); messages != nil {

		for _, message := range messages.Messages {
			slog.Info("message received", "message", *message.Body)
			creationRequest, err := toCreatePaymentRequest(*message.Body)
			if err != nil {
				slog.Error("error converting message to creation request", "error", err)
				continue
			}

			entity := dto.ToOrderEntity(creationRequest)
			order, err := r.kitchenUseCase.Creation(entity)
			if err != nil {
				slog.Error(fmt.Sprintf("error create order - %v", err))
				continue
			}

			slog.Info(fmt.Sprintf("create order %v request successfully", order.ID))
			r.awsSqsClient.Delete(message.ReceiptHandle, r.config.KitchenQueueUrl)
		}
	}
}

func toCreatePaymentRequest(jsonData string) (*dto.CreationRequest, error) {
	var creationRequest dto.CreationRequest
	err := json.Unmarshal([]byte(jsonData), &creationRequest)
	if err != nil {
		return nil, err
	}
	return &creationRequest, nil
}
