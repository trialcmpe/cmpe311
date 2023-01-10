package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBManager struct {
	cli         *mongo.Database
	collections *Collections
}

type Collections struct {
	UsersCollection   *mongo.Collection
	NotesCollection   *mongo.Collection
	CoursesCollection *mongo.Collection
}

func InitMongoDBConnection(connectionURI string, DBName string) (*MongoDBManager, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		return nil, err
	}

	c := client.Database(DBName)
	return &MongoDBManager{
		cli: c,
		collections: &Collections{
			UsersCollection:   c.Collection("Users"),
			NotesCollection:   c.Collection("Notes"),
			CoursesCollection: c.Collection("Courses"),
		},
	}, nil
}
