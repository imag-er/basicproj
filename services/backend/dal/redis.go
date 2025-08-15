package dal

import (
	"backend/config"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func initRedis() {
	redisConfig := config.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       0,
	})

	if err := client.Ping().Err(); err != nil {
		hlog.Fatalf("Failed to connect to Redis: %v", err)
	}

	RedisClient = client
}
