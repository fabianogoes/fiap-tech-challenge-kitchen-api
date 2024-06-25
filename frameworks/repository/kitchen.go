package repository

import (
	"context"
	"fmt"
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
	return &KitchenRepository{db, db.Collection("payments")}
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
	var order dbo.Order

	err := or.collection.FindOne(context.Background(), bson.M{"orderId": id}).Decode(&order)
	if err != nil {
		return nil, err
	}

	return order.ToOrderEntity(), nil
}

func (or *KitchenRepository) GetAll(status entities.OrderStatus) ([]*entities.Order, error) {
	if status == entities.OrderStatusUnknown {
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
	fmt.Println(order)
	return nil, nil
}
