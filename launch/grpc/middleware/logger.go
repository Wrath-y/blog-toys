package middleware

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
	grpcCtx "toys/infrastructure/common/context"
	"toys/infrastructure/common/errcode"
	"toys/infrastructure/util/logging"
	"toys/infrastructure/util/util/random"
)

const bodyLimitKB = 5000

func UnaryLogger() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if info.FullMethod == "/grpc.health.v1.Health/Check" {
			return handler(ctx, req)
		}

		start := time.Now()

		c := grpcCtx.GetContext(ctx)

		xRequestID := random.UUID()

		logger := logging.New()
		logger.SetRequestID(xRequestID)
		logger.Setv1(info.FullMethod)
		c.Logger = logger

		raw, _ := sonic.Marshal(req)

		rawKB := len(raw) / 1024 // => to KB
		if rawKB > bodyLimitKB {
			c.Logger.Info("接口请求与响应", string(raw[:1024]), nil)
			return nil, errcode.BlogBodyTooLarge.WithDetail(fmt.Sprintf("消息限制%dKB, 本消息%dKB", bodyLimitKB, rawKB))
		}

		defer func() {
			out, _ := sonic.Marshal(resp)
			md, _ := metadata.FromIncomingContext(c)

			request := map[string]any{
				"path":     info.FullMethod,
				"metadata": md,
				"body":     string(raw),
			}
			c.Logger.Info("接口请求与响应", request, string(out), logging.AttrOption{StartTime: &start})
		}()

		return handler(c, req)
	}
}
