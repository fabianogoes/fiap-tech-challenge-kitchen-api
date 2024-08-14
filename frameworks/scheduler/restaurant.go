package scheduler

import (
	"github.com/fabianogoes/fiap-kitchen/domain/ports"
	"github.com/robfig/cron"
	"log/slog"
)

func InitCronScheduler(restaurantReceiver ports.RestaurantReceiverPort) *cron.Cron {
	job := cron.New()

	_ = job.AddFunc("*/5 * * * *", restaurantReceiver.ReceiveOrder)

	job.Start()
	slog.Info("cron scheduler initialized")

	return job
}
