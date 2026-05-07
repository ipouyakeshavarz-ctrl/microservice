package interceptor

import (
	"context"
	"myapp/pkg/richerror"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func UnaryErrorInterceptor() grpc.UnaryServerInterceptor {

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		resp, err := handler(ctx, req)

		if err == nil {
			return resp, nil
		}

		if re, ok := err.(*richerror.RichError); ok {
			return nil, re.ToRPCError()
		}

		// اگر خطا RichError نبود، خطای معمولی gRPC بساز
		return nil, status.Convert(err).Err()
	}
}
