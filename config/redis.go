package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "go-micro.dev/v4/logger"
)

func initRedis() func() {
	client := redis.NewClient(&redis.Options{
		Addr:       Conf.Redis.Addr,
		Password:   Conf.Redis.Pwd,
		DB:         1,
		PoolSize:   30,
		MaxRetries: 5,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	KV = client

	return func() {
		if err := client.Close(); err != nil {
			log.Error(err)
		}
	}
}
