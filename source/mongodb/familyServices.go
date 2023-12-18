package mongodb

import (
	"context"

	"example.com/m/v2/libs/connection"
	"example.com/m/v2/source/models/mongodbModel"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const FamilyServicesCol = "familyServices"

type FamilyServices interface {
	Insert(ctx context.Context, item *mongodbModel.FamilyServices) (primitive.ObjectID, error)
	GetAll(ctx context.Context) (info []mongodbModel.FamilyServices, err error)
	GetByCode(ctx context.Context, code string) (mongodbModel.FamilyServices, bool, error)
}

type familyServicesDb struct {
	coll *mongo.Collection
}

func newFamilyServices(db *mongo.Database) FamilyServices {
	return &familyServicesDb{
		coll: connection.MongoInit(
			db,
			FamilyServicesCol,
			connection.MongoIndex{Keys: bson.D{{Key: mongodbModel.KeyIsRemoved, Value: 1}}},
			connection.MongoIndex{Keys: bson.D{{Key: mongodbModel.KeyCode, Value: 1}}},
		),
	}
}

func (ins *familyServicesDb) GetByCode(ctx context.Context, code string) (mongodbModel.FamilyServices, bool, error) {
	var (
		filter = bson.M{
			mongodbModel.KeyCode:      code,
			mongodbModel.KeyIsRemoved: false,
		}
		item = mongodbModel.FamilyServices{}
	)
	result := ins.coll.FindOne(ctx, filter)
	if result.Err() != nil {
		return mongodbModel.FamilyServices{}, false, result.Err()
	}
	if err := result.Decode(&item); err != nil {
		return mongodbModel.FamilyServices{}, false, errors.WithStack(err)
	}

	return item, true, nil
}

func (ins *familyServicesDb) Insert(ctx context.Context, item *mongodbModel.FamilyServices) (primitive.ObjectID, error) {

	result, err := ins.coll.InsertOne(ctx, item)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (db *familyServicesDb) GetAll(ctx context.Context) (info []mongodbModel.FamilyServices, err error) {
	info = make([]mongodbModel.FamilyServices, 0)

	filter := bson.M{
		mongodbModel.KeyIsRemoved: false,
	}
	//find all data
	opt := options.Find()
	opt.SetSort(bson.M{
		mongodbModel.KeyQuantifyIsUsed: -1,
	})
	cursor, err := db.coll.Find(ctx, filter, opt)
	// if not err
	if err == nil {
		err = cursor.All(ctx, &info)
	}
	defer cursor.Close(ctx)
	// if err
	return
}
