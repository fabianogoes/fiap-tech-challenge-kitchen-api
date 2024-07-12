package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/fabianogoes/fiap-kitchen/adapters/restaurant"
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

	rep := repository.NewKitchenRepository(db)
	restaurantAdapter := restaurant.NewClientAdapter(config)
	useCase := usecases.NewKitchenService(rep, &restaurantAdapter)
	handler := rest.NewKitchenHandler(useCase)

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
