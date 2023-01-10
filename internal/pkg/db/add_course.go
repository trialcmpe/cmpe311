package db

import (
	"context"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
)

func (d *MongoDBManager) AddCourse(ctx context.Context, course model.Course) error {
	if _, err := d.collections.CoursesCollection.InsertOne(ctx, course); err != nil {
		return err
	}

	return nil
}
