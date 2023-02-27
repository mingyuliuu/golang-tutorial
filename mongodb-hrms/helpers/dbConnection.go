package helpers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDBInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var MgDB MongoDBInstance

const dbName = "golang-hrms"
const mongoURI = "mongodb://localhost:27017" + dbName

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	MgDB = MongoDBInstance{
		Client:   client,
		Database: db,
	}

	return nil
}
