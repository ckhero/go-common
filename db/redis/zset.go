/**
 *@Description
 *@ClassName zset
 *@Date 2021/3/3 上午9:59
 *@Author ckhero
 */

package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Z struct {
	Score  float64
	Member interface{}
}

type ZRangeBy struct {
	Min, Max      string
	Offset, Count int64
}

func ZAdd(ctx context.Context, key string, vals ...*Z) error {
	key = buildKey(key)
	data := []*redis.Z{}
	for _, val := range vals {
		data = append(data, &redis.Z{
			Score:  val.Score,
			Member: val.Member,
		})
	}
	return GetGlobalRedisClient().ZAdd(ctx, key, data...).Err()
}

func ZAddEX(ctx context.Context, key string, expire time.Duration, vals ...*Z) error {
	if err := ZAdd(ctx, key, vals...); err != nil {
		return err
	}
	return Expire(ctx, key, expire)
}

func ZRangeByScoreWithScores(ctx context.Context, key string, query *ZRangeBy) ([]*Z, error) {
	key = buildKey(key)
	data, err := GetGlobalRedisClient().ZRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
		Min:    query.Min,
		Max:    query.Max,
		Offset: query.Offset,
		Count:  query.Count,
	}).Result()
	list := []*Z{}
	if err == redis.Nil {
		return list, nil
	}
	for _, v := range data {
		list = append(list, &Z{
			Score:  v.Score,
			Member: v.Member,
		})
	}
	return list, err
}
