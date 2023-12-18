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

const AccountCol = "account"

type Account interface {
	Insert(ctx context.Context, item *mongodbModel.Account) (primitive.ObjectID, error)
	GetByPhone(ctx context.Context, phone string) (mongodbModel.Account, bool, error)
	GetById(ctx context.Context, id primitive.ObjectID) (mongodbModel.Account, bool, error)

	GetListServiceProvider(ctx context.Context) (info []mongodbModel.Account, err error)
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
}

type accountDb struct {
	coll *mongo.Collection
}

func newAccount(db *mongo.Database) Account {
	return &accountDb{
		coll: connection.MongoInit(
			db,
			AccountCol,
		),
	}
}

func (ins *accountDb) Insert(ctx context.Context, item *mongodbModel.Account) (primitive.ObjectID, error) {

	result, err := ins.coll.InsertOne(ctx, item)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (ins *accountDb) GetByPhone(ctx context.Context, phone string) (mongodbModel.Account, bool, error) {
	var (
		filter = bson.M{
			mongodbModel.KeyPhone:     phone,
			mongodbModel.KeyIsRemoved: false,
		}
		item = mongodbModel.Account{}
	)
	result := ins.coll.FindOne(ctx, filter)
	if result.Err() != nil {
		return mongodbModel.Account{}, false, result.Err()
	}
	if err := result.Decode(&item); err != nil {
		return mongodbModel.Account{}, false, errors.WithStack(err)
	}

	return item, true, nil
}

func (ins *accountDb) GetById(ctx context.Context, id primitive.ObjectID) (mongodbModel.Account, bool, error) {
	var (
		filter = bson.M{
			mongodbModel.KeyId:        id,
			mongodbModel.KeyIsRemoved: false,
		}
		item = mongodbModel.Account{}
	)
	result := ins.coll.FindOne(ctx, filter)
	if result.Err() != nil {
		return mongodbModel.Account{}, false, result.Err()
	}
	if err := result.Decode(&item); err != nil {
		return mongodbModel.Account{}, false, errors.WithStack(err)
	}

	return item, true, nil
}

func (db *accountDb) GetListServiceProvider(ctx context.Context) (info []mongodbModel.Account, err error) {
	info = make([]mongodbModel.Account, 0)

	filter := bson.M{
		mongodbModel.KeyIsRemoved: false,
		mongodbModel.KeyRole:      2,
		mongodbModel.KeyStatus:    "",
	}

	cursor, err := db.coll.Find(ctx, filter)
	// if not err
	if err == nil {
		err = cursor.All(ctx, &info)
	}
	defer cursor.Close(ctx)
	// if err
	return
}

func (ins *accountDb) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {

	filter := bson.M{
		mongodbModel.KeyId: id,
	}
	// update data
	update := bson.M{"$set": bson.M{
		mongodbModel.KeyStatus: status,
	}}

	//update mongoDB
	_, err := ins.coll.UpdateOne(ctx, filter, update)

	return err
}
