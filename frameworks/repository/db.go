package repository

import (
	"context"
	"fmt"
	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InitDB(ctx context.Context, config *entities.Config) (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
	)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	fmt.Println("Successfully connected to MongoDB")

	return client.Database(config.DBName), nil
}
