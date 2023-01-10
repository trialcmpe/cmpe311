package db

import (
	"context"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
)

func (d *MongoDBManager) AddNote(ctx context.Context, note model.Note) error {
	if _, err := d.collections.NotesCollection.InsertOne(ctx, note); err != nil {
		return err
	}

	return nil
}
