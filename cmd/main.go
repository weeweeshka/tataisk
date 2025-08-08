package main

import (
	"github.com/weeweeshka/tataisk/internal/config"
	"github.com/weeweeshka/tataisk/pkg/lib/logger"
)

func main() {
	cfg := config.MustLoadConfig()

	logr := logger.SetupLogger()

	logr.Info("Config loaded")
	logr.Info("Logger initialized")

}
