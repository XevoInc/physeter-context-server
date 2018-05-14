package utils

import (
	"context"
	"path"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// Unary Interceptor
func UnaryServerInterceptor(log grpclog.LoggerV2) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var err error
		defer func(begin time.Time) {
			// infoからメソッド名を取得
			method := path.Base(info.FullMethod)
			took := time.Since(begin)
			if err != nil {
				log.Errorf("Failed RPC call: %s (%s), %s\n", method, took, err)
			} else {
				log.Infof("Succeeded RPC call: %s (%s)\n", method, took)
			}
		}(time.Now())

		reply, hErr := handler(ctx, req)
		if hErr != nil {
			err = hErr
		}
		return reply, err
	}
}

// Stream Iterceptor
func StreamServerInterceptor(log grpclog.LoggerV2) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		var err error
		defer func(begin time.Time) {
			method := path.Base(info.FullMethod)
			took := time.Since(begin)
			if err != nil {
				log.Errorf("Failed RPC call: %s (%s), %s\n", method, took, err)
			} else {
				log.Infof("Succeeded RPC call: %s (%s)\n", method, took)
			}
		}(time.Now())

		err = handler(srv, stream)
		return err
	}
}
