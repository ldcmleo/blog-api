package db

import (
	"context"
	"time"

	"github.com/ldcmleo/blog-api/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	uri, err := util.GetDBURI()
	if err != nil {
		return nil, err
	}

	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, connErr := mongo.Connect(ctx, clientOptions)
	if connErr != nil {
		return nil, connErr
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
