package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	client *mongo.Client
}

var ctx = context.Background()

func New(connstr string) (*MongoDb, error) {
	mongoOpts := options.Client().ApplyURI(connstr)
	client, err := mongo.Connect(ctx, mongoOpts)
	if err != nil {
		return nil, err
	}
	
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	
	return &MongoDb{
		client: client,
	}, nil
}