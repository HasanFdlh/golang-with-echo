package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASS"), // kosongkan "" jika tanpa password
		DB:       0,
	})

	// test koneksi
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("[SYSTEM ERROR] Redis connection failed:", err)
	}
	log.Println("[SYSTEM] Redis connected")
}
