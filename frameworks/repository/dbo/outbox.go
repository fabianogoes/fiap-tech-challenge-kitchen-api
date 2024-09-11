package dbo

import (
	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Outbox struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	OrderID     uint               `bson:"orderID"`
	MessageBody string             `bson:"messageBody"`
	CreatedAt   time.Time          `bson:"created_at"`
	QueueUrl    string             `bson:"queueUrl"`
}

func (o *Outbox) ToEntity() *entities.Outbox {
	return &entities.Outbox{
		ID:          o.ID.Hex(),
		CreatedAt:   o.CreatedAt,
		MessageBody: o.MessageBody,
		QueueUrl:    o.QueueUrl,
	}
}

func ToOutboxDBO(orderID uint, messageBody string, queueUrl string) Outbox {
	return Outbox{
		OrderID:     orderID,
		MessageBody: messageBody,
		CreatedAt:   time.Now(),
		QueueUrl:    queueUrl,
	}
}
