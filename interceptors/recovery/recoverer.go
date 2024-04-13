package recovery

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RPCRecover struct{}

func (rec *RPCRecover) CustomRecoveryFuncWithContext(ctx context.Context, p interface{}) error {
	logger := zerolog.Ctx(ctx)
	logger.Error().Err(fmt.Errorf("%+v", p)).Stack().Msg("recover from panic")
	return status.Error(codes.Internal, "uncommon internal panic error")
}
