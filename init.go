package zlogwrap

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func New(conf ...Config) ZerologWrapper {
	// set default config
	cfg := configDefault(conf...)
	return cfg
}

type Config struct {

	// Optional. Default: false
	Hidden bool

	// Optional. Default: ""
	ServiceName string

	// Optional. Default: false
	TerminateOnCriticalLevel bool

	// Optional. Default: &fiber.Ctx{}
	Context *fiber.Ctx

	// Optional. Default: log.Logger
	Logger zerolog.Logger
}

var ConfigDefault = Config{
	Hidden:                   false,
	ServiceName:              "",
	TerminateOnCriticalLevel: false,
	Context:                  nil,
	Logger:                   log.Logger,
}

func configDefault(config ...Config) ZerologWrapper {
	// Return default config if nothing provided
	if len(config) < 1 {
		return &ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	// Note: cfg.Hidden: it's false by default.
	// Note: cfg.ServiceName: it's "" by default.
	// Note: cfg.TerminateOnCriticalLevel: it's false by default.
	// Note: cfg.Context: it's nil by default.

	if reflect.ValueOf(cfg.Logger).IsZero() {
		cfg.Logger = log.Logger
	}

	return &cfg
}
