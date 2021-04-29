/**
 *@Description
 *@ClassName context
 *@Date 2020/11/29 10:53 上午
 *@Author ckhero
 */

package context

import (
	"context"
	sysConst "github.com/ckhero/go-common/constant/sys"
	"github.com/gin-gonic/gin"
)

// ContextWithSpan 返回context
func ContextWithSpan(c *gin.Context) (ctx context.Context, ok bool) {

	v, exist := c.Get(sysConst.SysContextTracerKey)
	if !exist {
		ctx, ok = context.TODO(), false
		return
	}
	ctx, ok = v.(context.Context)
	return
}

