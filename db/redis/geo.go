/**
 *@Description
 *@ClassName geo
 *@Date 2021/3/2 下午4:22
 *@Author ckhero
 */

package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	xerrors "github.com/pkg/errors"
	"strconv"
)

type GeoLocation struct {
	Name                      string
	NameUint64                uint64
	Longitude, Latitude, Dist float64
	GeoHash                   int64
}

type GeoRadiusQuery struct {
	Radius float64
	// Can be m, km, ft, or mi. Default is km.
	Unit        string
	WithCoord   bool
	WithDist    bool
	WithGeoHash bool
	Count       int
	// Can be ASC or DESC. Default is no sort order.
	Sort      string
	Store     string
	StoreDist string
}

func GeoAdd(ctx context.Context, key string, geoLocation ...*GeoLocation) error {
	key = buildKey(key)
	list := []*redis.GeoLocation{}
	for _, v := range geoLocation {
		list = append(list, &redis.GeoLocation{
			Name:      v.Name,
			Longitude: v.Longitude,
			Latitude:  v.Latitude,
		})
	}
	return GetGlobalRedisClient().GeoAdd(ctx, key, list...).Err()
}

func GeoRadius(ctx context.Context, key string, lng, lat float64, query *GeoRadiusQuery) ([]*GeoLocation, error) {
	key = buildKey(key)
	query01 := &redis.GeoRadiusQuery{}
	_ = util.DeepCopy(query, query01)
	res, err := GetGlobalRedisClient().GeoRadius(ctx, key, lng, lat, query01).Result()
	if err == redis.Nil {
		return []*GeoLocation{}, errors_v2.NotFound("geo not exists!", "key [%s]", key)
	}
	if err != nil {
		return []*GeoLocation{}, xerrors.Wrapf(err, "key [%s]", key)
	}
	return convert2OwnGeoLocation(res), nil
}

func GeoRadiusUint64(ctx context.Context, key string, lng, lat float64, query *GeoRadiusQuery) ([]*GeoLocation, error) {
	key = buildKey(key)
	query01 := &redis.GeoRadiusQuery{}
	_ = util.DeepCopy(query, query01)
	res, err := GetGlobalRedisClient().GeoRadius(ctx, key, lng, lat, query01).Result()
	if err == redis.Nil {
		return []*GeoLocation{}, errors_v2.NotFound("geo not exists!", "key [%s]", key)
	}
	if err != nil {
		return []*GeoLocation{}, xerrors.Wrapf(err, "key [%s]", key)
	}
	return convert2OwnGeoLocationUint64(res), nil
}

func GeoDist(ctx context.Context, key string, member1, member2, unit string) (float64, error) {
	key = buildKey(key)
	res, err := GetGlobalRedisClient().GeoDist(ctx, key, member1, member2, unit).Result()
	if err == redis.Nil {
		return 0, nil
	}
	return res, err
}

func GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) ([]*GeoLocation, error) {
	key = buildKey(key)
	res, err := GetGlobalRedisClient().GeoRadiusByMember(ctx, key, member, query).Result()
	if err == redis.Nil {
		return []*GeoLocation{}, nil
	}
	if err != nil {
		return []*GeoLocation{}, nil
	}
	return convert2OwnGeoLocation(res), nil
}

func convert2OwnGeoLocation(data []redis.GeoLocation) []*GeoLocation {
	list := []*GeoLocation{}
	for _, v := range data {
		list = append(list, &GeoLocation{
			Name:      v.Name,
			Longitude: v.Longitude,
			Latitude:  v.Latitude,
			Dist:      v.Dist,
			GeoHash:   v.GeoHash,
		})
	}
	return list
}

func convert2OwnGeoLocationUint64(data []redis.GeoLocation) []*GeoLocation {
	list := []*GeoLocation{}
	for _, v := range data {
		nameUint64, _ := strconv.ParseUint(v.Name, 10, 64)
		list = append(list, &GeoLocation{
			NameUint64: nameUint64,
			Longitude:  v.Longitude,
			Latitude:   v.Latitude,
			Dist:       v.Dist,
			GeoHash:    v.GeoHash,
		})
	}
	return list
}
