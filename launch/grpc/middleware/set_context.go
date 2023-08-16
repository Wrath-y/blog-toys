package middleware

import (
	"context"
	"google.golang.org/grpc"
	grpcCtx "toys/infrastructure/common/context"
)

func UnaryContext() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		c := grpcCtx.NewContext(ctx)

		return handler(c, req)
	}
}
