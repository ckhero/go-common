/**
 *@Description
 *@ClassName gin
 *@Date 2020/11/30 9:38 下午
 *@Author ckhero
 */

package gin

import (
	sysConst "github.com/ckhero/go-common/constant/sys"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)
// jaeger
func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		carrier := opentracing.HTTPHeadersCarrier(c.Request.Header)
		spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, carrier)
		span := opentracing.GlobalTracer().StartSpan(c.Request.URL.Path, opentracing.ChildOf(spanCtx))
		defer span.Finish()

		//http 和 rpc之间打通
		c.Set(sysConst.SysContextTracerKey, opentracing.ContextWithSpan(c.Request.Context(), span))

		if sp, ok := span.Context().(jaeger.SpanContext); ok {
			//将trace-id返回到http header中
			c.Writer.Header().Set(sysConst.SysHeaderTrace, sp.TraceID().String())
		}
		c.Next()
	}
}
