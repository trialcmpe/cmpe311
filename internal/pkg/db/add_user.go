package db

import (
	"context"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
)

func (d *MongoDBManager) AddUser(ctx context.Context, user model.User) error {
	if _, err := d.collections.UsersCollection.InsertOne(ctx, user); err != nil {
		return err
	}

	return nil
}
