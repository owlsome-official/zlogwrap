package zlogwrap

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

const ( // Todo: It's should be a config
	ServiceNameTag         = "service"
	RequestIDHeaderKeyTag  = "transaction-id"
	RequestIDContextKeyTag = "transaction-id"
	URLTag                 = "url"
)

type ZerologWrapper interface {
	SetField(key string, anything interface{}) ZerologWrapper // Set field in logs
	Debug(anything ...interface{})                            // level 0
	Info(anything ...interface{})                             // level 1
	Warn(anything ...interface{})                             // level 2
	Error(anything ...interface{})                            // level 3
	Fatal(anything ...interface{})                            // level 4
	Panic(anything ...interface{})                            // level 5
	GetLogEvent(zerolog.Level) *zerolog.Event                 // With Caller (file:line)
}

func (c Config) SetField(key string, anything interface{}) ZerologWrapper {
	if key == "" {
		return &c
	}
	switch v := anything.(type) {
	case bool:
		c.Logger = c.Logger.With().Bool(key, v).Logger()
	case int:
		c.Logger = c.Logger.With().Int(key, v).Logger()
	case int64:
		c.Logger = c.Logger.With().Int64(key, v).Logger()
	case float64:
		c.Logger = c.Logger.With().Float64(key, v).Logger()
	case []byte:
		c.Logger = c.Logger.With().RawJSON(key, v).Logger()
	case []int:
		c.Logger = c.Logger.With().Ints(key, v).Logger()
	case []int64:
		c.Logger = c.Logger.With().Ints64(key, v).Logger()
	case []float64:
		c.Logger = c.Logger.With().Floats64(key, v).Logger()
	case []string:
		c.Logger = c.Logger.With().Strs(key, v).Logger()
	default:
		c.Logger = c.Logger.With().Interface(key, v).Logger()
	}
	return &c
}

func (c *Config) createLogTemplate(zLevel zerolog.Level) *zerolog.Event {
	var logTemplate *zerolog.Event

	logTemplate = c.Logger.WithLevel(zLevel)

	if c.ServiceName != "" {
		logTemplate = logTemplate.Str(ServiceNameTag, c.ServiceName)
	}

	if c.Context != nil {
		// RequestID: Header style
		if reqID := string(c.Context.Response().Header.Peek(RequestIDHeaderKeyTag)); reqID != "" {
			logTemplate = logTemplate.Str(strings.ReplaceAll(RequestIDHeaderKeyTag, "-", "_"), reqID)
		}

		// RequestID: ContextKey style
		if reqID := c.Context.Locals(RequestIDContextKeyTag); reqID != nil {
			logTemplate = logTemplate.Str(strings.ReplaceAll(RequestIDContextKeyTag, "-", "_"), reqID.(string))
		}

		if url := c.Context.OriginalURL(); url != "" {
			logTemplate = logTemplate.Str(URLTag, url)
		}
	}

	return logTemplate
}

func (c *Config) Debug(anything ...interface{}) {
	if c.Hidden {
		return
	}

	logString := toString(anything...)

	logTemplate := c.createLogTemplate(zerolog.DebugLevel)
	logTemplate.Msgf("%v", logString)
}

func (c *Config) Info(anything ...interface{}) {
	if c.Hidden {
		return
	}

	logString := toString(anything...)

	logTemplate := c.createLogTemplate(zerolog.InfoLevel)
	logTemplate.Msgf("%v", logString)
}

func (c *Config) Warn(anything ...interface{}) {
	if c.Hidden {
		return
	}

	logString := toString(anything...)

	logTemplate := c.createLogTemplate(zerolog.WarnLevel)
	logTemplate.Msgf("%v", logString)
}

func (c *Config) Error(anything ...interface{}) {
	if c.Hidden {
		return
	}

	logString := toString(anything...)

	logTemplate := c.createLogTemplate(zerolog.ErrorLevel)
	logTemplate.Msgf("%v", logString)
}

func (c *Config) Fatal(anything ...interface{}) {
	if c.Hidden {
		return
	}

	logString := toString(anything...)

	logTemplate := c.createLogTemplate(zerolog.FatalLevel)
	logTemplate.Msgf("%v", logString)
	if c.TerminateOnCriticalLevel {
		os.Exit(1)
	}
}

func (c *Config) Panic(anything ...interface{}) {
	if c.Hidden {
		return
	}

	logString := toString(anything...)

	logTemplate := c.createLogTemplate(zerolog.PanicLevel)
	logTemplate.Msgf("%v", logString)
	if c.TerminateOnCriticalLevel {
		panic(logString)
	}
}

func (c *Config) GetLogEvent(level zerolog.Level) *zerolog.Event {
	return c.Logger.WithLevel(level)
}
