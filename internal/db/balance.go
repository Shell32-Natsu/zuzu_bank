package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Balance(ctx context.Context, id int64) (int64, error) {
	var result bson.M
	err := balanceCollection().FindOne(ctx, bson.D{{Key: "user_id", Value: id}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, createAccount(ctx, id)
		}
		return 0, fmt.Errorf("failed to query balance for %d: %s", id, err)
	}
	if r, ok := result["balance"].(int64); !ok {
		return 0, fmt.Errorf("invalid balance value from database")
	} else {
		return r, nil
	}
}

func createAccount(ctx context.Context, id int64) error {
	_, err := balanceCollection().InsertOne(ctx,
		bson.D{
			{
				Key:   "user_id",
				Value: id,
			},
			{
				Key:   "balance",
				Value: int64(0),
			},
		})
	return err
}
