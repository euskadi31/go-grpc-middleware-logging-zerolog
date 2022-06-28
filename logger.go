package zerolog

import (
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
)

// Compatibility check.
var _ logging.Logger = (*Logger)(nil)

// Logger is a zerolog logging adapter compatible with logging middlewares.
type Logger struct {
	zerolog.Logger

	fieldPrefix string
}

type Option func(l *Logger)

func WithFieldPrefix(prefix string) Option {
	return func(l *Logger) {
		l.fieldPrefix = prefix
	}
}

// InterceptorLogger is a zerolog.Logger to Logger adapter.
func InterceptorLogger(logger zerolog.Logger, opts ...Option) *Logger {
	l := &Logger{
		Logger: logger,
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}

// Log implements the logging.Logger interface.
func (l *Logger) Log(lvl logging.Level, msg string) {
	switch lvl {
	case logging.DEBUG:
		l.Debug().Msg(msg)
	case logging.INFO:
		l.Info().Msg(msg)
	case logging.WARNING:
		l.Warn().Msg(msg)
	case logging.ERROR:
		l.Error().Msg(msg)
	default:
		l.Error().Msgf("zerolog: unknown level %s using error", lvl)
		l.Error().Msg(msg)
	}
}

// With implements the logging.Logger interface.
func (l Logger) With(fields ...string) logging.Logger {
	vals := make(map[string]interface{}, len(fields)/2)
	for i := 0; i < len(fields); i += 2 {
		vals[l.formatField(fields[i])] = fields[i+1]
	}

	return InterceptorLogger(l.Logger.With().Fields(vals).Logger())
}

func (l Logger) formatField(field string) string {
	if l.fieldPrefix == "" {
		return field
	}

	return strings.TrimRight(l.fieldPrefix, ".") + "." + field
}
