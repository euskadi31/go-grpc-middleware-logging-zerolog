package zerolog

import (
	"bytes"
	"testing"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	logger := zerolog.New(buf)

	l := InterceptorLogger(logger)

	l.With("foo", "bar").Log(logging.DEBUG, "test")
	l.With("foo", "bar").Log(logging.INFO, "test")
	l.With("foo", "bar").Log(logging.WARNING, "test")
	l.With("foo", "bar").Log(logging.ERROR, "test")
	l.With("foo", "bar").Log(logging.Level("bad"), "test")

	assert.Equal(t, "{\"level\":\"debug\",\"foo\":\"bar\",\"message\":\"test\"}\n{\"level\":\"info\",\"foo\":\"bar\",\"message\":\"test\"}\n{\"level\":\"warn\",\"foo\":\"bar\",\"message\":\"test\"}\n{\"level\":\"error\",\"foo\":\"bar\",\"message\":\"test\"}\n{\"level\":\"error\",\"foo\":\"bar\",\"message\":\"zerolog: unknown level bad using error\"}\n{\"level\":\"error\",\"foo\":\"bar\",\"message\":\"test\"}\n", string(buf.Bytes()))
}

func TestLogWithPrefix(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	logger := zerolog.New(buf)

	l := InterceptorLogger(logger, WithFieldPrefix("labels."))

	l.With("foo", "bar").Log(logging.DEBUG, "test")

	assert.Equal(t, "{\"level\":\"debug\",\"labels.foo\":\"bar\",\"message\":\"test\"}\n", string(buf.Bytes()))
}

func TestLogWithPrefixAndSubSubWith(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	logger := zerolog.New(buf)

	l := InterceptorLogger(logger, WithFieldPrefix("labels."))

	l.With("foo-1", "bar1").With("foo-2", "bar2").Log(logging.DEBUG, "test")

	assert.Equal(t, "{\"level\":\"debug\",\"labels.foo-1\":\"bar1\",\"labels.foo-2\":\"bar2\",\"message\":\"test\"}\n", string(buf.Bytes()))
}
