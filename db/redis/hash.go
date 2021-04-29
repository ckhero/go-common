/**
 *@Description
 *@ClassName hash
 *@Date 2021/4/28 上午10:54
 *@Author ckhero
 */

package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func HSet(ctx context.Context, key string, vals ...interface{}) (int64, error) {
	key = buildKey(key)
	return GetGlobalRedisClient().HSet(ctx, key, vals...).Result()
}

func HGet(ctx context.Context, key string, field string) (string, error) {
	key = buildKey(key)
	res, err := GetGlobalRedisClient().HGet(ctx, key, field).Result()
	if err == redis.Nil {
		return "", nil
	}
	return res, err
}