/**
 *@Description
 *@ClassName set
 *@Date 2021/3/3 上午10:55
 *@Author ckhero
 */

package redis

import (
	"context"
	"time"
)

func SAdd(ctx context.Context, key string, vals ...interface{}) (int64, error) {
	key = buildKey(key)
	return GetGlobalRedisClient().SAdd(ctx, key, vals...).Result()
}

// 移除集合中一个或多个成员
func SREM(ctx context.Context, key string, vals ...interface{}) (int64, error) {
	key = buildKey(key)
	return GetGlobalRedisClient().SRem(ctx, key, vals...).Result()
}

func SAddEX(ctx context.Context, key string, expire time.Duration, vals ...interface{}) (int64, error) {
	res, err := SAdd(ctx, key, vals)
	if err != nil {
		return 0, err
	}
	if err := Expire(ctx, key, expire); err != nil {
		return 0, err
	}

	return res, nil
}

func SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	key = buildKey(key)
	return GetGlobalRedisClient().SIsMember(ctx, key, member).Result()
}
