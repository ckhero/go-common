/**
 *@Description
 *@ClassName web
 *@Date 2020/11/28 10:17 下午
 *@Author ckhero
 */

package web

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	ginCommon "youmi-micro-cluster/src/common/middleware/gin"
	validator2 "youmi-micro-cluster/src/common/validator"
)

type RegisterRouter func(*gin.Engine)


type RegisterValidator func(*validator.Validate) error

func Run(errc chan error, rr RegisterRouter, rv RegisterValidator) {

	//注册路由
	engine := gin.New()
	engine.Use(
		ginCommon.TraceMiddleware(),
		ginCommon.RecoverMiddleware(),
		ginCommon.LoggerMiddleware(),
		)
	//注册顺序很重要，勿轻易调整，跨域请求已经转移到nginx层处理，此处不再处理cors
	rr(engine)

	// 注册验证器
	if err := validator2.InitValidatorTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
	//  注册自定义翻译器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := rv(v)
		if err != nil {
			logger.GetLogger(context.TODO()).Error("注册验证器失败")
			return
		}
	}
	//启动
	port := fmt.Sprintf(":%d", config.AppConfig.HTTP.Port)
	if err := engine.Run(port); err != nil {
		errc <- err
	}
}
