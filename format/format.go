package format

import (
	sysConst "github.com/ckhero/go-common/constant/sys"
	"github.com/ckhero/go-common/logger"
	"github.com/ckhero/go-common/util/context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

/**
 * 成功
 */
func Success(ctx *gin.Context, data interface{}) {
	d := gin.H{
		"code": 0,
		"msg":  "成功",
		"data": data,
	}
	ctx.Set(sysConst.SysResponseDataKey, d)
	ctx.JSON(http.StatusOK, d)
}

/**
 * 失败
 */
func Fail(ctx *gin.Context, errs ...interface{}) {
	c, _ := context.ContextWithSpan(ctx)

	code := 10000
	msg := "系统错误"
	for _, e := range errs {
		switch e.(type) {
		case string:
			msg = e.(string)
		case error:
			logger.GetLogger(c).Errorf("%v", errors.Cause(e.(error)))
		}


		data := gin.H{
			"code": code,
			"msg":  msg,
			"data": gin.H{},
		}
		ctx.Set(sysConst.SysResponseDataKey, data)
		ctx.JSON(http.StatusOK, data)
	}
}
