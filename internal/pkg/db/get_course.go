package db

import (
	"context"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (d *MongoDBManager) GetCourse(ctx context.Context, courseName string) (*model.Course, error) {
	filter := bson.D{{Key: "name", Value: courseName}}
	var course model.Course
	err := d.collections.CoursesCollection.FindOne(ctx, filter).Decode(&course)
	if err != nil {
		return nil, err
	}

	return &course, err
}

func (d *MongoDBManager) GetCourses(ctx context.Context) (*[]model.Course, error) {
	var course []model.Course
	cursor, err := d.collections.CoursesCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &course); err != nil {
		return nil, err
	}

	return &course, err
}
