/**
 *@Description
 *@ClassName context
 *@Date 2020/11/29 10:53 上午
 *@Author ckhero
 */

package context

import (
	"context"
	"github.com/gin-gonic/gin"
	"youmi-micro-cluster/src/common/constant"
	ginUtil "youmi-micro-cluster/src/common/gin"
)

// ContextWithSpan 返回context
func ContextWithSpan(c *gin.Context) (ctx context.Context, ok bool) {

	v, exist := c.Get(constant.SysContextTracerKey)
	if !exist {
		ctx, ok = context.TODO(), false
		return
	}
	ctx, ok = v.(context.Context)
	ctx = context.WithValue(ctx, constant.HeadShopId, ginUtil.GetShopIdFromGin(c))
	return
}

func GetShopIdFromCtx(ctx context.Context) uint64 {
	return ctx.Value(constant.HeadShopId).(uint64)
}

