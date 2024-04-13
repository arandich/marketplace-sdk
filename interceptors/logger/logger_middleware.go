package logger_middleware

import (
	"context"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func NewUnaryLoggerInterceptor(ctx context.Context) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger := zerolog.Ctx(ctx).With().Str("method", info.FullMethod).Logger()
		logger.Info().Msg("received request")

		resp, err := handler(ctx, req)
		if err == nil {
			logger.Error().Err(err).Msg("error handled when processing request")

			return resp, nil
		}

		logger.Info().Msg("request successfully processed")

		return nil, err
	}
}
