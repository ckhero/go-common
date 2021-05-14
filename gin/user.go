/**
 *@Description
 *@ClassName user
 *@Date 2020/12/28 下午9:39
 *@Author ckhero
 */

package gin

import (
	"github.com/ckhero/go-common/auth"
	"github.com/ckhero/go-common/constant/sys"
	"github.com/ckhero/go-common/errors"
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
		userId, data, err := auth.ResolveJWTToken(token, secretKey, logger.GetLogger(ctx))
		c.Set(sys.SysKeyUserId, userId)
		openId, ok := data["openId"]
		if !ok || openId == "" || userId == 0 {
			c.Abort()
			format.Fail(c, errors.Unauthorized("user", "尚未登录", "尚未登录"))
			return
		}
		c.Set(sys.SysKeyOpenId, openId)
		if err != nil {
			c.Abort()
			format.Fail(c, errors.Unauthorized("user", "尚未登录", "尚未登录"))
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

func GetUserId(c *gin.Context) uint64 {
	return  c.GetUint64(sys.SysKeyUserId)
}

func GetOpenId(c *gin.Context) string {
	return  c.GetString(sys.SysKeyOpenId)
}