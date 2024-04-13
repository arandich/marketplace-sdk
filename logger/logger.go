package logger

import (
	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"log"
	"os"
)

func NewLogger(levelStr string) *zerolog.Logger {
	logLevel, err := zerolog.ParseLevel(levelStr)
	if err != nil {
		logLevel = zerolog.InfoLevel
	}

	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "message"
	zerolog.ErrorStackFieldName = "stacktrace"
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	zerolog.InterfaceMarshalFunc = json.Marshal
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(logLevel)
	zerolog.ErrorHandler = func(err error) {
		log.Fatal(err)
	}

	sampler := zerolog.LevelSampler{
		TraceSampler: &zerolog.BasicSampler{N: 2},
		DebugSampler: &zerolog.BasicSampler{N: 4},
		InfoSampler:  &zerolog.BasicSampler{N: 2},
	}

	logger := zerolog.New(os.Stderr).
		Sample(sampler).
		With().
		Timestamp().
		Logger()

	return &logger
}
