/**
 *@Description
 *@ClassName logger
 *@Date 2020/12/29 上午11:09
 *@Author ckhero
 */

package gin

import (
	"encoding/json"
	sysConst "github.com/ckhero/go-common/constant/sys"
	"github.com/ckhero/go-common/logger"
	"github.com/ckhero/go-common/util/context"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由，注意提取get中的参数
		reqUri := c.Request.RequestURI
		reqUriArray := strings.Split(reqUri, "?")

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 请求返回值
		var response []byte
		if responseData, ok := c.Get(sysConst.SysResponseDataKey); ok {
			if responseDataMap, ok := responseData.(gin.H); ok {
				response, _ = json.Marshal(responseDataMap)
			}
		}

		data := map[string]interface{}{
			"status": statusCode,
			"spend": float64(endTime.Sub(startTime)) / 1e6,
			"clientIp": clientIP,
			"method": reqMethod,
			"uri": reqUriArray[0],
			"ua": c.Request.UserAgent(),
			"response": string(response),
		}

		ctx, _ := context.ContextWithSpan(c)

		// 日志记录
		logger.GetLogger(ctx).Info(data)
	}
}

