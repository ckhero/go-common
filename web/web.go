/**
 *@Description
 *@ClassName web
 *@Date 2020/11/28 10:17 下午
 *@Author ckhero
 */

package web

import (
	"fmt"
	"github.com/ckhero/go-common/config"
	ginCommon "github.com/ckhero/go-common/gin"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterRouter func(*gin.Engine)


type RegisterValidator func(*validator.Validate) error

func Run(errc chan error, rr RegisterRouter, rv RegisterValidator) {

	//注册路由
	engine := gin.New()
	engine.Use(
		ginCommon.TraceMiddleware(),
		)
	//注册顺序很重要，勿轻易调整，跨域请求已经转移到nginx层处理，此处不再处理cors
	rr(engine)

	//// 注册验证器
	//if err := validator2.InitValidatorTrans("zh"); err != nil {
	//	fmt.Printf("init trans failed, err:%v\n", err)
	//	return
	//}
	////  注册自定义翻译器
	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	err := rv(v)
	//	if err != nil {
	//		logger.GetLogger(context.TODO()).Error("注册验证器失败")
	//		return
	//	}
	//}
	//启动
	fmt.Println(config.GetGlobalCfg())
	port := fmt.Sprintf(":%d", config.GetGlobalCfg().HTTP.Port)
	if err := engine.Run(port); err != nil {
		errc <- err
	}
}
