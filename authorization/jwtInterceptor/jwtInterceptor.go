package jwtInterceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ChainUnaryInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	auth := metadata.ValueFromIncomingContext(ctx, "authorization")
	if len(auth) != 0 {
		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", auth[0])
	}
	return invoker(ctx, method, req, reply, cc, opts...)
}
