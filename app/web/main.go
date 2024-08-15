package main

import (
	"context"
	"fmt"
	"github.com/fabianogoes/fiap-kitchen/adapters/messaging"
	"github.com/fabianogoes/fiap-kitchen/frameworks/scheduler"
	"log/slog"
	"os"

	"github.com/fabianogoes/fiap-kitchen/domain/usecases"
	"github.com/fabianogoes/fiap-kitchen/frameworks/repository"

	"github.com/fabianogoes/fiap-kitchen/domain/entities"

	"github.com/fabianogoes/fiap-kitchen/frameworks/rest"
)

func init() {
	fmt.Println("Initializing...")

	var logHandler *slog.JSONHandler

	config, _ := entities.NewConfig()
	if config.Environment == "production" {
		logHandler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}

func main() {
	fmt.Println("Starting web server...")

	ctx := context.Background()
	var err error

	config, err := entities.NewConfig()
	if err != nil {
		panic(err)
	}
	db, err := repository.InitDB(ctx, config)
	if err != nil {
		panic(err)
	}

	sqsClient := messaging.NewAWSSQSClient(config)
	rep := repository.NewKitchenRepository(db)
	outboxRepository := repository.NewOutboxRepository(db)
	restaurantPublisher := messaging.NewRestaurantPublisher(config, sqsClient, outboxRepository)
	useCase := usecases.NewKitchenService(rep, restaurantPublisher)
	handler := rest.NewKitchenHandler(useCase)

	restaurantMessaging := messaging.NewRestaurantReceiver(useCase, config, sqsClient)
	outboxRetry := messaging.NewOutboxRetry(sqsClient, outboxRepository)
	cron := scheduler.InitCronScheduler(restaurantMessaging, outboxRetry)
	defer cron.Stop()

	router, err := rest.NewRouter(handler)
	if err != nil {
		panic(err)
	}

	fmt.Println("DB connected")
	fmt.Println(db)

	err = router.Run(config.AppPort)
	if err != nil {
		panic(err)
	}
}
