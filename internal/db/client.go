package db

import (
	"context"
	"fmt"
	"log"

	"github.com/Shell32-Natsu/zuzu_bank/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName                    = "zuzu_bank"
	balanceCollectionName     = "balance"
	transactionCollectionName = "transaction"
)

var client *mongo.Client

func Init(ctx context.Context) error {
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.MongoDBConnectionString()))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %s", err)
	}
	log.Printf("connected to database")

	if err := checkAndCreateUniqueIndex(ctx); err != nil {
		return err
	}

	return nil
}

func checkAndCreateUniqueIndex(ctx context.Context) error {
	const expectedIndexName = "user_id_1"
	indexSpecs, err := balanceCollection().Indexes().ListSpecifications(ctx, options.ListIndexes())
	if err != nil {
		return fmt.Errorf("failed to list index specifications: %s", err)
	}
	for _, is := range indexSpecs {
		if expectedIndexName == is.Name {
			log.Printf("index %q has existed", is.Name)
			return nil
		}
	}

	indexName, err := balanceCollection().Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "user_id", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return fmt.Errorf("failed to create unique index: %s", err)
	}
	log.Printf("created unique index %q", indexName)
	return nil
}

func balanceCollection() *mongo.Collection {
	return client.Database(dbName).Collection(balanceCollectionName)
}

func transactionCollection() *mongo.Collection {
	return client.Database(dbName).Collection(transactionCollectionName)
}
