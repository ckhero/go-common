/**
 *@Description
 *@ClassName funk
 *@Date 2021/1/26 上午10:38
 *@Author ckhero
 */

package funk

import (
	"context"
	"github.com/thoas/go-funk"
	"youmi-micro-cluster/src/common/logger"
)

func GetUint64s(src interface{}, path string) []uint64 {
	defer func() {
		if e := recover(); e != nil {
			logger.GetLogger(context.TODO()).Errorf("get uint64s fail %s", e)
		}
	}()
	if src == nil {
		return []uint64{}
	}
	res := funk.Get(src, path)
	if res == nil {
		return []uint64{}
	}
	return res.([]uint64)
}
