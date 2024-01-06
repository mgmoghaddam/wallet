package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"wallet/internal/config"
)

func SetupLogger() error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	lvl, err := zerolog.ParseLevel(config.LogLevel())
	if err != nil {
		return fmt.Errorf("failed to pars level: %v", err)
	}
	zerolog.SetGlobalLevel(lvl)
	return nil
}
