package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewLogger() zerolog.Logger {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	return log.Logger
}
