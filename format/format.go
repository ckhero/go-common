package format

import (
	"github.com/gin-gonic/gin"
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
	ctx.JSON(http.StatusOK, d)
}

/**
 * 失败
 */
func Fail(ctx *gin.Context, errs ...interface{}) {
	code := 10000
	msg := "系统错误"
	for _, e := range errs {
		switch e.(type) {
		case string:
			msg = e.(string)
		}

		data := gin.H{
			"code": code,
			"msg":  msg,
			"data": gin.H{},
		}
		ctx.JSON(http.StatusOK, data)
	}
}
