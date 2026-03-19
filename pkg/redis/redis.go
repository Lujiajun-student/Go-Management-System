package redis

import (
	"Go-Management-System/common/config"
	"context"

	"github.com/go-redis/redis/v8"
)

// RedisClient 得到的Redis客户端
var (
	RedisClient *redis.Client
)

// InitRedis 初始化Redis
func InitRedis() {

	// 获取config里的Redis
	redisConfig := config.Config.Redis
	// 创建Redis客户端
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	// 测试连接
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}
