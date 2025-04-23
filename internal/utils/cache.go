package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var RedisClient *redis.Client

func InitRedis(addr string) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test connection
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
}

func SetCache(key string, value interface{}, expiration time.Duration) error {
	return RedisClient.Set(context.Background(), key, value, expiration).Err()
}

func GetCache(key string) (string, error) {
	return RedisClient.Get(context.Background(), key).Result()
}

func DeleteCache(key string) error {
	return RedisClient.Del(context.Background(), key).Err()
}
