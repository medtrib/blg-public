package redisHelper

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"math"
	"time"
)

// GetRedisCacheKey 根据参数获取 Redis key
func GetRedisCacheKeyByParams(module string, params map[string]interface{}) string {
	cacheKey := module + ":"
	if len(params) == 0 {
		return cacheKey
	}
	jsonParams, _ := json.Marshal(params)
	paramsStr := string(jsonParams)
	cacheKey += paramsStr
	return cacheKey
}

// SaveRedisCacheWithoutExpiration 缓存信息（无有效期）
func SaveRedisCacheWithoutExpiration(client redis.Cmdable, key, value string) error {
	return client.Set(key, value, 0).Err()
}

// SaveRedisCacheWithExpiration 缓存信息，并设置有效期
func SaveRedisCacheWithExpiration(client redis.Cmdable, key, value string, expiration time.Duration) error {
	return client.Set(key, value, expiration).Err()
}

// GetRedisCache 获取信息缓存
func GetRedisCache(client redis.Cmdable, key string) string {
	return client.Get(key).Val()
}

// DeleteRedisCache 删除信息缓存
func DeleteRedisCache(client redis.Cmdable, key string) {
	client.Del(key)
}

// BatchDeleteRedisCache 批量删除信息缓存
func BatchDeleteRedisCache(client redis.Cmdable, prefix string) {
	keys, _ := client.Scan(0, prefix+"*", math.MaxInt64).Val()
	client.Del(keys...)
}
