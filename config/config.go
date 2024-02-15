package config

import (
	"time"

	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
	"github.com/silvergama/transations/pkg/logger"
	"go.uber.org/zap"
)

type app struct {
	Name string `env:"APP_NAME" envDefault:"transaction"`
}

type serverHTTP struct {
	Port         int           `env:"SERVER_PORT" envDefault:"8080"`
	WriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT" envDefault:"15s"`
	ReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT" envDefault:"15s"`
}

type database struct {
	Host string `env:"DATABASE_HOST" envDefault:"localhost"`
	Port string `env:"DATABASE_PORT" envDefault:"5432"`
	User string `env:"DATABASE_USER,required"`
	Pwd  string `env:"DATABASE_PWD,required"`
	Base string `env:"DATABASE_BASE" envDefault:"transaction"`
}

type Config struct {
	App        app
	ServerHTTP serverHTTP
	Database   database
}

func ReadProperties() Config {
	logger.Info("Loading environments...")

	var cfg Config

	opts := env.Options{
		RequiredIfNoDef: true,
	}

	if err := env.Parse(&cfg, opts); err != nil {
		logger.Panic(
			"failed to read properties",
			zap.Error(err),
		)
	}

	return cfg
}
