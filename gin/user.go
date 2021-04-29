/**
 *@Description
 *@ClassName user
 *@Date 2020/12/28 下午9:39
 *@Author ckhero
 */

package gin

import (
	"github.com/ckhero/go-common/auth"
	"github.com/ckhero/go-common/format"
	"github.com/ckhero/go-common/logger"
	"github.com/ckhero/go-common/util/context"
	"github.com/ckhero/go-common/util/param"
	"github.com/gin-gonic/gin"
)

/* gin用户jwt认证中间件 */
func UserJwtAuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, _ := context.ContextWithSpan(c)
		token := param.GetToken(c)
		_, _, err := auth.ResolveJWTToken(token, secretKey, logger.GetLogger(ctx))
		if err != nil {
			c.Abort()
			format.Fail(c, err)
			return
		}
		c.Next()
	}
}

func NeedLogonMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Abort()
		format.Fail(c, "用户尚未登陆")
		return
	}
}