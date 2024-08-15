package repository

import (
	"context"
	"github.com/fabianogoes/fiap-kitchen/domain/entities"
	"github.com/fabianogoes/fiap-kitchen/frameworks/repository/dbo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OutboxRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewOutboxRepository(db *mongo.Database) *OutboxRepository {
	return &OutboxRepository{db, db.Collection("outboxes")}
}

func (or *OutboxRepository) GetOutboxById(id string) (*entities.Outbox, error) {
	var outbox dbo.Outbox

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	err = or.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&outbox)
	if err != nil {
		return nil, err
	}

	return outbox.ToEntity(), nil
}

func (or *OutboxRepository) CreateOutbox(orderID uint, messageBody string, queueUrl string) (*entities.Outbox, error) {
	paymentCreate := dbo.ToOutboxDBO(orderID, messageBody, queueUrl)

	res, err := or.collection.InsertOne(context.Background(), paymentCreate)
	if err != nil {
		return nil, err
	}

	id := res.InsertedID.(primitive.ObjectID)
	outboxResponse, err := or.GetOutboxById(id.Hex())
	if err != nil {
		return nil, err
	}

	return outboxResponse, nil
}

func (or *OutboxRepository) DeleteOutbox(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	_, err = or.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (or *OutboxRepository) GetAll() ([]*entities.Outbox, error) {
	ctx := context.Background()
	cursor, err := or.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*dbo.Outbox
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	var list []*entities.Outbox
	for _, result := range results {
		list = append(list, result.ToEntity())
	}

	return list, nil
}
