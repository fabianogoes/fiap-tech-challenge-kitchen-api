package entities

import "time"

type Outbox struct {
	ID          string
	OrderID     uint
	CreatedAt   time.Time
	MessageBody string
	QueueUrl    string
}
