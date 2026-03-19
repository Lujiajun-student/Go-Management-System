package util

import (
	"Go-Management-System/common/constant"
	"Go-Management-System/pkg/redis"
	"context"
	"log"
	"time"
)

// Redis 存取验证码，这里的ctx为无限制的上下文，以后需要通过context.WithTimeout来设置超时时间，避免redis无限等待
var ctx = context.Background()

type RedisStore struct {
}

// Set 存验证码
func (r RedisStore) Set(id string, value string) error {
	key := constant.LOGIN_CODE + id
	// 通过Redis客户端存储键值对
	err := redis.RedisClient.Set(ctx, key, value, time.Minute*5).Err()
	if err != nil {
		log.Panicln(err.Error())
		return err
	}
	return nil
}

// Get 获取验证码
func (r RedisStore) Get(id string, clear bool) string {
	key := constant.LOGIN_CODE + id
	val, err := redis.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return val
}

// Verify 校验验证码
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
