package main

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func configureLogger() *zerolog.Logger {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	output := (io.Writer)(os.Stderr)
	if DevMode {
		output = zerolog.ConsoleWriter{Out: os.Stderr}
	}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return &logger
}
