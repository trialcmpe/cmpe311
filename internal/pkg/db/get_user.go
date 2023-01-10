package db

import (
	"context"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (d *MongoDBManager) GetUser(ctx context.Context, email string) (*model.User, error) {
	filter := bson.D{{Key: "email", Value: email}}
	var user model.User
	err := d.collections.UsersCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, err
}
