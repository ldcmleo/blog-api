package database

import (
	"context"

	"github.com/ldcmleo/blog-api/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDatabaseClient() (*mongo.Client, error) {
	uri, envError := util.GetMongoDBCredentials()
	if envError != nil {
		return nil, envError
	}

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	errPing := client.Ping(context.TODO(), nil)
	if errPing != nil {
		return nil, errPing
	}

	return client, nil
}
