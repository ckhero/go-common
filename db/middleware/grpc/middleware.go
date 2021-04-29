/**
 *@Description
 *@ClassName middleware
 *@Date 2020/11/26 12:03 上午
 *@Author ckhero
 */

package grpc

import (
	"context"
	"fmt"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime/debug"
	"youmi-micro-cluster/src/common/logger"
)

var DefaultMiddlewareOpts = []grpc.ServerOption{

	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpctransport.Interceptor,
		RecoveryInterceptorMiddleware,
		JaegerServerMiddleware(),
	)),
}

func RecoveryInterceptorMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
			fmt.Println(err)
		}
	}()

	return handler(ctx, req)
}

func LoggerMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	logger.GetLogger(ctx).WithField("req", req).WithField("method", info).Info("request ! ")
	resp, err = handler(ctx, req)
	logger.GetLoggerWithBody(ctx, resp).Info("response")
	return
}

