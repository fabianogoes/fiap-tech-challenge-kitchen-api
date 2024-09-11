package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"github.com/fabianogoes/fiap-kitchen/frameworks/repository/dbo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type KitchenRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewKitchenRepository(db *mongo.Database) *KitchenRepository {
	return &KitchenRepository{db, db.Collection("orders")}
}

func (or *KitchenRepository) Create(order *entities.Order) (*entities.Order, error) {
	orderCreate := dbo.ToOrderDBO(order)

	_, err := or.collection.InsertOne(context.Background(), orderCreate)
	if err != nil {
		return nil, err
	}

	orderResponse, err := or.GetById(order.ID)
	if err != nil {
		return nil, err
	}

	return orderResponse, nil
}

func (or *KitchenRepository) GetById(id uint) (*entities.Order, error) {
	log.Default().Printf("GetById orderID: %d \n", id)
	var order dbo.Order

	err := or.collection.FindOne(context.Background(), bson.M{"orderId": int(id)}).Decode(&order)
	if err != nil {
		return nil, err
	}

	return order.ToOrderEntity(), nil
}

func (or *KitchenRepository) GetAll(status entities.OrderStatus) ([]*entities.Order, error) {
	if status == entities.OrderStatusKitchenUnknown {
		return nil, fmt.Errorf("invalid status")
	}

	cursor, err := or.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	orders := make([]*entities.Order, 0)
	for cursor.Next(context.Background()) {
		var order dbo.Order
		err := cursor.Decode(&order)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order.ToOrderEntity())
	}

	return orders, nil
}

func (or *KitchenRepository) UpdateStatus(order *entities.Order) (*entities.Order, error) {
	update := bson.M{"$set": bson.M{
		"status": order.Status.ToString(),
	}}

	one, err := or.collection.UpdateOne(context.Background(), bson.M{"orderId": order.ID}, update)
	fmt.Printf("Update one %v\n", one)
	if err != nil {
		fmt.Printf("Error updating order %v status %v \n", order.ID, order.Status.ToString())
		return nil, err
	}

	orderResponse, err := or.GetById(order.ID)
	if err != nil {
		return nil, err
	}

	return orderResponse, nil
}
