/**
 *@Description
 *@ClassName base
 *@Date 2021/3/3 上午9:15
 *@Author ckhero
 */

package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	xerrors "github.com/pkg/errors"
	"time"
)

func buildKey(key string) string {
	return getPrefix() + key
}

func Set(ctx context.Context, key string, val interface{}, expire time.Duration) error {
	key = buildKey(key)
	return GetGlobalRedisClient().Set(ctx, key, val, expire).Err()
}

func SetNX(ctx context.Context, key string, val interface{}, expire time.Duration) (bool, error) {
	key = buildKey(key)
	return GetGlobalRedisClient().SetNX(ctx, key, val, expire).Result()
}

func Get(ctx context.Context, key string) (res interface{}, err error) {
	key = buildKey(key)
	res, err = GetGlobalRedisClient().Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	return
}

func GetInt64(ctx context.Context, key string) (res int64, err error) {
	key = buildKey(key)

	res, err = GetGlobalRedisClient().Get(ctx, key).Int64()
	if err == redis.Nil {
		return 0, err
	}

	return
}

func GetUint64(ctx context.Context, key string) (res uint64, err error) {
	key = buildKey(key)

	res, err = GetGlobalRedisClient().Get(ctx, key).Uint64()
	if err == redis.Nil {
		return 0, err
	}

	return
}

func GetFloat64(ctx context.Context, key string) (res float64, err error) {
	key = buildKey(key)

	res, err = GetGlobalRedisClient().Get(ctx, key).Float64()
	if err == redis.Nil {
		return 0, err
	}

	return
}

func GetString(ctx context.Context, key string) string {
	key = buildKey(key)
	return GetGlobalRedisClient().Get(ctx, key).Val()
}

func Expire(ctx context.Context, key string, expire time.Duration) error {
	key = buildKey(key)
	return GetGlobalRedisClient().Expire(ctx, key, expire).Err()
}

func Exists(ctx context.Context, key string) (bool, error) {
	key = buildKey(key)
	res, err := GetGlobalRedisClient().Exists(ctx, key).Result()
	if err != nil {
		return false, xerrors.Wrapf(err, "not exists key [%s]", key)
	}
	return res > 0, err
}

func Del(ctx context.Context, keys ...string) (bool, error) {
	for k, _ := range keys {
		keys[k] = buildKey(keys[k])
	}
	res, err := GetGlobalRedisClient().Del(ctx, keys...).Result()
	if err != nil {
		return false, xerrors.Wrapf(err, "dels keys fail [%s]", keys)
	}
	return res > 0, err
}

func Incr(ctx context.Context, key string) (int64, error) {
	key = buildKey(key)
	return GetGlobalRedisClient().Incr(ctx, key).Result()
}

func IncrEX(ctx context.Context, key string, expire time.Duration) (int64, error) {
	key = buildKey(key)
	res, err := GetGlobalRedisClient().Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	err = Expire(ctx, key, expire)
	return res, err
}

func Decr(ctx context.Context, key string) (int64, error) {
	key = buildKey(key)
	return GetGlobalRedisClient().Decr(ctx, key).Result()
}

func DecrEX(ctx context.Context, key string, expire time.Duration) (int64, error) {
	key = buildKey(key)
	res, err := Decr(ctx, key)
	if err != nil {
		return 0, err
	}
	err = Expire(ctx, key, expire)
	return res, err
}

func IncrBy(ctx context.Context, key string, val int64) (int64, error) {
	key = buildKey(key)
	return GetGlobalRedisClient().IncrBy(ctx, key, val).Result()
}

func IncrByEX(ctx context.Context, key string, val int64, expire time.Duration) (int64, error) {
	key = buildKey(key)
	res, err := IncrBy(ctx, key, val)
	if err != nil {
		return 0, err
	}
	err = Expire(ctx, key, expire)
	return res, err
}

func DecrBy(ctx context.Context, key string, val int64) (int64, error) {
	key = buildKey(key)
	return GetGlobalRedisClient().DecrBy(ctx, key, val).Result()
}

func DecrByEX(ctx context.Context, key string, val int64, expire time.Duration) (int64, error) {
	key = buildKey(key)
	res, err := DecrBy(ctx, key, val)
	if err != nil {
		return 0, err
	}
	err = Expire(ctx, key, expire)
	return res, err
}
