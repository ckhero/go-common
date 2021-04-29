/**
 *@Description
 *@ClassName redis
 *@Date 2020/11/23 4:59 下午
 *@Author ckhero
 */

package plugin

import (
	"fmt"
	"github.com/ckhero/go-common/config"
	"github.com/ckhero/go-common/db/redis"
)

type Redis struct {
}

func NewPluginRedis() *Redis {
	return &Redis{
	}
}

func (r *Redis) InitPlugin() error {
	for k, v := range config.GetRedisCfg() {
		redis.ConnectRedisV1(k, v)
	}
	return nil
}

func (r *Redis) Release() {
	_ = redis.GetGlobalRedisClient().Close()
	fmt.Println("redis release")

}