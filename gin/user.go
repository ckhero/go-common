/**
 *@Description
 *@ClassName user
 *@Date 2020/12/28 下午9:39
 *@Author ckhero
 */

package gin

import (
	"github.com/ckhero/go-common/auth"
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
		userId, info, err := auth.ResolveJWTToken(token, secretKey, logger.GetLogger(ctx))
		if err != nil {
			c.Abort()
			format.Fail(c, err)
			return
		}

		c.Set(constant.CtxKeyUserId, userId)
		c.Set(constant.CtxKeyUserInfo, info)
		c.Next()
	}
}

func NeedLogonMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		if auth2.GetUserId(c) > 0 {
			c.Next()
			return
		}

		c.Abort()
		format.Fail(c, errors.ErrorUserLogin)
		return
	}
}