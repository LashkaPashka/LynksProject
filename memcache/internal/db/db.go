package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Db struct {
	RedisDb *redis.Client
}

var ctx = context.Background()

func New() (*Db, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "mysecretpassword",
		DB: 0,
	})

	if res := redisClient.Ping(ctx); res.Err() != nil {
		return nil, res.Err()
	}

	//defer redisClient.Close()
	
	return &Db{
		RedisDb: redisClient,
	}, nil
}