package scheduler

import (
	"github.com/fabianogoes/fiap-kitchen/adapters/messaging"
	"github.com/fabianogoes/fiap-kitchen/domain/ports"
	"github.com/robfig/cron"
	"log/slog"
)

func InitCronScheduler(
	restaurantReceiver ports.RestaurantReceiverPort,
	outboxRetry *messaging.OutboxRetry,
) *cron.Cron {
	job := cron.New()

	_ = job.AddFunc("*/5 * * * *", restaurantReceiver.ReceiveOrder)
	_ = job.AddFunc("*/30 * * * *", outboxRetry.Retry)

	job.Start()
	slog.Info("cron scheduler initialized")

	return job
}
