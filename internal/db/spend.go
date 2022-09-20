package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func Spend(ctx context.Context, user_id, amount int64, desc string) (int64, int64, error) {
	currentBalance, err := Balance(ctx, user_id)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get current balance: %s", err)
	}
	if currentBalance < amount {
		return 0, 0, fmt.Errorf("余额不足哦，只有%d", currentBalance)
	}
	filter := bson.D{{Key: "user_id", Value: user_id}}
	update := bson.D{{
		Key: "$set",
		Value: bson.D{
			{
				Key:   "balance",
				Value: currentBalance - amount,
			},
		},
	}}

	var result bson.M
	err = balanceCollection().FindOneAndUpdate(ctx, filter, update).Decode(&result)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to update balance: %s", err)
	}

	if _, ok := result["balance"].(int64); !ok {
		return 0, 0, fmt.Errorf("invalid balance value from database")
	} else {
		return currentBalance, currentBalance - amount, nil
	}
}
