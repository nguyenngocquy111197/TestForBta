package mongodb

import (
	"context"

	"example.com/m/v2/libs/connection"
	"example.com/m/v2/source/models/mongodbModel"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const TransactionCol = "transaction"

type Transaction interface {
	Insert(ctx context.Context, item *mongodbModel.Transaction) (primitive.ObjectID, error)
	UpdateIdServiceProvider(ctx context.Context, transactionId, serviceProviderId primitive.ObjectID, status string) error
	UpdateStatus(ctx context.Context, transactionId primitive.ObjectID, status string) error
	GetByTransactionId(ctx context.Context, transactionId primitive.ObjectID) (mongodbModel.Transaction, bool, error)
}

type transactionDb struct {
	coll *mongo.Collection
}

func newTransaction(db *mongo.Database) Transaction {
	return &transactionDb{
		coll: connection.MongoInit(
			db,
			TransactionCol,
		),
	}
}

func (ins *transactionDb) Insert(ctx context.Context, item *mongodbModel.Transaction) (primitive.ObjectID, error) {

	result, err := ins.coll.InsertOne(ctx, item)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (ins *transactionDb) UpdateIdServiceProvider(ctx context.Context, transactionId, serviceProviderId primitive.ObjectID, status string) error {

	filter := bson.M{
		mongodbModel.KeyId: transactionId,
	}
	// update data
	update := bson.M{"$set": bson.M{
		mongodbModel.KeyIdServiceProvider: serviceProviderId,
		mongodbModel.KeyStatus:            status,
	}}

	//update mongoDB
	_, err := ins.coll.UpdateOne(ctx, filter, update)

	return err
}

func (ins *transactionDb) UpdateStatus(ctx context.Context, transactionId primitive.ObjectID, status string) error {

	filter := bson.M{
		mongodbModel.KeyId: transactionId,
	}
	// update data
	update := bson.M{"$set": bson.M{
		mongodbModel.KeyStatus: status,
	}}

	//update mongoDB
	_, err := ins.coll.UpdateOne(ctx, filter, update)

	return err
}

func (ins *transactionDb) GetByTransactionId(ctx context.Context, transactionId primitive.ObjectID) (mongodbModel.Transaction, bool, error) {
	var (
		filter = bson.M{
			mongodbModel.KeyId: transactionId,
		}
		item = mongodbModel.Transaction{}
	)
	result := ins.coll.FindOne(ctx, filter)
	if result.Err() != nil {
		return mongodbModel.Transaction{}, false, result.Err()
	}
	if err := result.Decode(&item); err != nil {
		return mongodbModel.Transaction{}, false, errors.WithStack(err)
	}

	return item, true, nil
}
