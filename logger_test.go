package zerolog

import (
	"bytes"
	"context"
	"testing"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	ctx := context.Background()
	buf := bytes.NewBuffer([]byte{})

	logger := zerolog.New(buf)

	l := InterceptorLogger(logger)

	l.With("foo", "bar").Log(ctx, logging.LevelDebug, "test")
	l.With("foo", "bar").Log(ctx, logging.LevelInfo, "test")
	l.With("foo", "bar").Log(ctx, logging.LevelWarn, "test")
	l.With("foo", "bar").Log(ctx, logging.LevelError, "test")
	l.With("foo", "bar").Log(ctx, logging.Level(9), "test")

	assert.Equal(t, "{\"level\":\"debug\",\"foo\":\"bar\",\"message\":\"test\"}\n{\"level\":\"info\",\"foo\":\"bar\",\"message\":\"test\"}\n{\"level\":\"warn\",\"foo\":\"bar\",\"message\":\"test\"}\n{\"level\":\"error\",\"foo\":\"bar\",\"message\":\"test\"}\n{\"level\":\"error\",\"foo\":\"bar\",\"message\":\"zerolog: unknown level 9 using error\"}\n{\"level\":\"error\",\"foo\":\"bar\",\"message\":\"test\"}\n", string(buf.Bytes()))
}

func TestLogWithPrefix(t *testing.T) {
	ctx := context.Background()
	buf := bytes.NewBuffer([]byte{})

	logger := zerolog.New(buf)

	l := InterceptorLogger(logger, WithFieldPrefix("labels."))

	l.With("foo", "bar").Log(ctx, logging.LevelDebug, "test")

	assert.Equal(t, "{\"level\":\"debug\",\"labels.foo\":\"bar\",\"message\":\"test\"}\n", string(buf.Bytes()))
}

func TestLogWithPrefixAndSubSubWith(t *testing.T) {
	ctx := context.Background()
	buf := bytes.NewBuffer([]byte{})

	logger := zerolog.New(buf)

	l := InterceptorLogger(logger, WithFieldPrefix("labels."))

	l.With("foo-1", "bar1").With("foo-2", "bar2").Log(ctx, logging.LevelDebug, "test")

	assert.Equal(t, "{\"level\":\"debug\",\"labels.foo-1\":\"bar1\",\"labels.foo-2\":\"bar2\",\"message\":\"test\"}\n", string(buf.Bytes()))
}
