/**
 *@Description
 *@ClassName database
 *@Date 2021/5/5 上午11:12
 *@Author ckhero
 */

package db

import (
	"context"
	"github.com/ckhero/go-common/db/mysql"
	redisSelf "github.com/ckhero/go-common/db/redis"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Database struct {
	RDB func(ctx context.Context) *gorm.DB
	Redis func() *redis.Client
}

func NewDatabase() *Database {
	return &Database{
		Redis: redisSelf.LazyRedisClient(),
		RDB: mysql.LazyGetDefaultDB(),
	}
}
