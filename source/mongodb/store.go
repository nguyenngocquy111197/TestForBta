package mongodb

import (
	"context"

	"log"

	"time"

	"example.com/m/v2/libs/connection"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
Store database
*/
type Store struct {
	FamilyServices FamilyServices
	Account        Account
	Transaction    Transaction
	db             *mongo.Database
}

// Drop :
func (s *Store) Drop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := s.db.Drop(ctx)
	return err
}

/*
New create a new store
*/
func New(uri, dbname string, timeout time.Duration) *Store {
	db, err := connection.MongoConnect(uri, dbname, timeout)
	if err != nil {
		log.Fatal(err)
	}
	return &Store{
		FamilyServices: newFamilyServices(db),
		Account:        newAccount(db),
		Transaction:    newTransaction(db),
		//
		db: db,
	}
}
