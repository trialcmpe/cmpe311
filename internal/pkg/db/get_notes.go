package db

import (
	"context"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (d *MongoDBManager) GetNotes(ctx context.Context, course string) (*[]model.Note, error) {
	filter := bson.D{{Key: "course", Value: course}}

	cursor, err := d.collections.NotesCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var notes []model.Note
	if err = cursor.All(ctx, &notes); err != nil {
		return nil, err
	}

	return &notes, nil
}
