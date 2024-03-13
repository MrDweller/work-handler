package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database
var Work *mongo.Collection
var Worker *mongo.Collection

func InitDatabase(dbConnectionString string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	Client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbConnectionString))
	if err != nil {
		return err
	}
	Database = Client.Database("WorkHandler")
	Work = Database.Collection("Work")
	Worker = Database.Collection("Worker")

	return nil
}
