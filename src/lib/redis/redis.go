//go:build !unit

package redis

import (
	"HRSystem/config"
	"context"
	"fmt"
	"log"
	"time"

	goredis "github.com/go-redis/redis/v8"
)

type (
	CacheClient struct {
		*goredis.Client
	}
)

var (
	DB *CacheClient
)

func Init() *goredis.Client {
	DB = new(CacheClient)
	DB.Client = goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Config.RedisHost, config.Config.RedisPort),
		Password: config.Config.RedisPassword,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := DB.Client.Ping(ctx).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Redis connection established.")
	return DB.Client
}
